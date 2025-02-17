"use client";

import { ScrollArea } from "@/components/ui/scroll-area";
import { ChevronRight, ChevronLeft, User, Bot, FunctionSquare } from "lucide-react";
import type { Team, Component, ToolConfig, AssistantAgentConfig, UserProxyAgentConfig } from "@/types/datamodel";
import { Button } from "@/components/ui/button";
import { Badge } from "@/components/ui/badge";
import { Tooltip, TooltipContent, TooltipProvider, TooltipTrigger } from "@/components/ui/tooltip";

interface AgentDetailsPanelProps {
  selectedTeam: Team | null;
  isOpen: boolean;
  onToggle: () => void;
}

export function AgentDetailsPanel({ selectedTeam, isOpen, onToggle }: AgentDetailsPanelProps) {
  const renderAgentTools = (tools: Component<ToolConfig>[] = []) => {
    if (tools.length === 0) {
      return <div className="text-sm text-white/40 italic">No tools available</div>;
    }

    return (
      <ul className="mt-4 flex flex-col gap-2">
        {tools.map((tool, index) => (
          <li key={tool.label + "_" + index}>
            <TooltipProvider>
              <Tooltip>
                <TooltipTrigger asChild>
                  <span className="hover:text-white/60 transition-colors">{tool.label}</span>
                </TooltipTrigger>
                <TooltipContent>
                  <p>{tool.description}</p>
                </TooltipContent>
              </Tooltip>
            </TooltipProvider>
          </li>
        ))}
      </ul>
    );
  };

  return (
    <div
      className={`fixed top-0 right-0 h-screen transition-all duration-300 ease-in-out 
          bg-[#2A2A2A] border-l border-t border-b border-[#3A3A3A] 
          ${isOpen ? "w-96" : "w-12"}`}
    >
      <div className="h-full flex flex-col text-white">
        <div className="p-4 flex items-center gap-2 border-b border-[#3A3A3A]">
          {isOpen && <h1 className="text-sm font-semibold flex-1">Agent</h1>}
          <Button variant="ghost" size="icon" onClick={onToggle} className="h-8 w-8 hover:bg-[#3A3A3A] text-white hover:text-white transition-colors">
            {isOpen ? <ChevronRight className="h-4 w-4" /> : <ChevronLeft className="h-4 w-4" />}
          </Button>
          x
        </div>

        <div className={`h-full ${isOpen ? "opacity-100" : "opacity-0"} transition-opacity duration-200`}>
          <div className="h-full flex flex-col text-white">
            {!selectedTeam ? (
              <div className="flex items-center justify-center flex-1 text-white/50 p-6">No agent selected</div>
            ) : (
              <ScrollArea className="flex-1 px-6 py-6">
                <div className="space-y-4">
                  {/* filter out the planner agent we add automatically */}
                  {selectedTeam.component?.config.participants
                    .filter((a) => !a.label?.startsWith("kagent_"))
                    .map((participant, index) => {
                      // TODO: Move the providers to consts + add support for other agent types
                      if (participant.provider === "autogen_agentchat.agents.AssistantAgent") {
                        const assistantAgent = participant.config as AssistantAgentConfig;
                        return (
                          <div key={index} className="text-sm text-white/50">
                            <div className="flex items-center justify-between">
                              <h5 className="text-white font-semibold text-lg flex items-center gap-2">
                                <Bot className="h-6 w-6 text-violet-500" />
                                {assistantAgent.name}
                              </h5>
                              <TooltipProvider>
                                <Tooltip>
                                  <TooltipTrigger asChild>
                                    <Badge variant="outline" className="text-xs text-white/50">
                                      {assistantAgent.model_client.config.model}
                                    </Badge>
                                  </TooltipTrigger>
                                  <TooltipContent>
                                    <p>Model agent is using</p>
                                  </TooltipContent>
                                </Tooltip>
                              </TooltipProvider>
                            </div>
                            <p className="mt-4 text-base">{assistantAgent.description}</p>

                            <div className="mt-8">
                              <h6 className="text-base font-medium text-white/70 flex items-center gap-2 mb-4">
                                <FunctionSquare className="h-6 w-6" />
                                Available Tools
                              </h6>
                              {renderAgentTools(assistantAgent.tools)}
                            </div>
                          </div>
                        );
                      }

                      const userProxyAgent = participant.config as UserProxyAgentConfig;
                      return (
                        <div key={index} className="text-sm text-white/50 p-4 rounded-md bg-[#1A1A1A] border border-[#3A3A3A] hover:bg-[#222] transition-colors">
                          <div className="flex items-center justify-between">
                            <h5 className="text-white font-semibold flex items-center gap-2">
                              <User className="h-4 w-4 text-blue-500" />
                              {participant.label}
                            </h5>
                            <Badge variant="outline" className="text-xs text-white/50">
                              {participant.label}
                            </Badge>
                          </div>

                          <p className="mt-2">{userProxyAgent.description}</p>
                        </div>
                      );
                    })}
                </div>
              </ScrollArea>
            )}
          </div>
        </div>
      </div>
    </div>
  );
}
