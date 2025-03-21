package client

import "fmt"

func (c *Client) ListSessions(userID string) ([]*Session, error) {
	var sessions []*Session
	err := c.doRequest("GET", fmt.Sprintf("/sessions/?user_id=%s", userID), nil, &sessions)
	return sessions, err
}

func (c *Client) CreateSession(session *CreateSession) (*Session, error) {
	var result Session
	err := c.doRequest("POST", "/sessions/", session, &result)
	return &result, err
}

func (c *Client) GetSession(sessionID int, userID string) (*Session, error) {
	var session Session
	err := c.doRequest("GET", fmt.Sprintf("/sessions/%d?user_id=%s", sessionID, userID), nil, &session)
	return &session, err
}

func (c *Client) DeleteSession(sessionID int, userID string) error {
	return c.doRequest("DELETE", fmt.Sprintf("/sessions/%d?user_id=%s", sessionID, userID), nil, nil)
}

func (c *Client) ListSessionRuns(sessionID int, userID string) ([]*Run, error) {
	var runs SessionRuns
	err := c.doRequest("GET", fmt.Sprintf("/sessions/%d/runs/?user_id=%s", sessionID, userID), nil, &runs)
	if err != nil {
		return nil, err
	}

	var result []*Run
	for _, run := range runs.Runs {
		result = append(result, &run)
	}
	return result, err
}
