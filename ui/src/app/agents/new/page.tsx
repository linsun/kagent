"use client";
import React, { useEffect, useState } from "react";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";
import { Textarea } from "@/components/ui/textarea";
import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Alert, AlertDescription } from "@/components/ui/alert";
import { isIdentifier } from "@/lib/utils";
import { Bot, ArrowLeft, Loader2, FunctionSquare } from "lucide-react";
import { Model } from "@/lib/types";
import { createTeamConfig, transformToAgentConfig } from "@/lib/agents";
import { SystemPromptSection } from "@/components/create/SystemPromptSection";
import { ModelSelectionSection } from "@/components/create/ModelSelectionSection";
import { AVAILABLE_MODELS } from "@/lib/data";
import { ToolsSection } from "@/components/create/ToolsSection";
import { createTeam } from "@/app/actions/teams";
import { useRouter } from "next/navigation";
import { getBuiltInTools } from "@/app/actions/tools";
import { Component, ToolConfig } from "@/types/datamodel";

interface ValidationErrors {
  name?: string;
  description?: string;
  systemPrompt?: string;
  model?: string;
  knowledgeSources?: string;
  tools?: string;
}

export default function NewAgentPage() {
  const router = useRouter();

  // Basic form state
  const [name, setName] = useState("");
  const [description, setDescription] = useState("");
  const [systemPrompt, setSystemPrompt] = useState("");

  // Default to the first model
  const [selectedModel, setSelectedModel] = useState<Model>(AVAILABLE_MODELS[0]);
  const [allTools, setAllTools] = useState<Component<ToolConfig>[]>([]);

  // Tools state
  const [selectedTools, setSelectedTools] = useState<Component<ToolConfig>[]>([]);

  // Overall form state
  const [isSubmitting, setIsSubmitting] = useState(false);
  const [errors, setErrors] = useState<ValidationErrors>({});
  const [generalError, setGeneralError] = useState("");

  useEffect(() => {
    const fetchAllTools = async () => {
      // Fetch all tools
      const response = await getBuiltInTools();
      if (response.success && response.data) {
        setAllTools(response.data);
      }
    };
    fetchAllTools();
  }, []);

  const validateForm = () => {
    const newErrors: ValidationErrors = {};

    if (!name.trim()) {
      newErrors.name = "Agent name is required";
    } else if (!isIdentifier(name)) {
      newErrors.name = "Agent name can't contain spaces or special characters";
    } else if (name.length > 50) {
      newErrors.name = "Agent name must be less than 50 characters";
    }

    if (!description.trim()) {
      newErrors.description = "Description is required";
    }

    if (!systemPrompt.trim()) {
      newErrors.systemPrompt = "Agent instructions are required";
    }

    if (!selectedModel) {
      newErrors.model = "Please select a model";
    }

    setErrors(newErrors);
    return Object.keys(newErrors).length === 0;
  };

  const handleCreateAgent = async () => {
    if (!validateForm()) {
      return;
    }

    try {
      setIsSubmitting(true);
      setGeneralError("");

      const agentData = {
        name,
        description,
        system_prompt: systemPrompt,
        model: selectedModel,
        tools: selectedTools,
      };

      const agentConfig = transformToAgentConfig(agentData);
      const teamConfig = await createTeamConfig(agentConfig);

      const result = await createTeam(teamConfig);

      if (!result.success) {
        throw new Error("Failed to create team");
      }

      if (!result.data || !result.data.id) {
        router.push(`/`);
      }

      router.push(`/agents/${result.data?.id}/chat`);
    } catch (error) {
      console.error("Error creating agent:", error);
      setGeneralError("Failed to create agent. Please try again.");
    } finally {
      setIsSubmitting(false);
    }
  };

  return (
    <div className="min-h-screen bg-[#1A1A1A] text-white p-8">
      <Button variant="ghost" className="mb-8" onClick={() => window.history.back()} disabled={isSubmitting}>
        <ArrowLeft className="mr-2 h-4 w-4" />
        Back
      </Button>

      <div className="max-w-4xl mx-auto">
        <h1 className="text-2xl font-bold mb-8">Create New Agent</h1>

        {generalError && (
          <Alert variant="destructive" className="mb-6">
            <AlertDescription>{generalError}</AlertDescription>
          </Alert>
        )}

        <div className="space-y-6">
          <Card className="bg-[#2A2A2A] border-[#3A3A3A]">
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-white">
                <Bot className="h-5 w-5 text-violet-500" />
                Basic Information
              </CardTitle>
            </CardHeader>
            <CardContent className="space-y-4">
              <div>
                <label className="text-sm text-white/70 mb-2 block">Agent Name</label>
                <Input
                  value={name}
                  onChange={(e) => setName(e.target.value)}
                  className={`bg-[#1A1A1A] border-[#3A3A3A] text-white ${errors.name ? "border-red-500" : ""}`}
                  placeholder="Enter agent name..."
                  disabled={isSubmitting}
                />
                {errors.name && <p className="text-red-500 text-sm mt-1">{errors.name}</p>}
              </div>

              <div>
                <label className="text-sm text-white/70 mb-2 block">Description</label>
                <Textarea
                  value={description}
                  onChange={(e) => setDescription(e.target.value)}
                  className={`bg-[#1A1A1A] border-[#3A3A3A] text-white min-h-[100px] ${errors.description ? "border-red-500" : ""}`}
                  placeholder="Describe your agent's purpose..."
                  disabled={isSubmitting}
                />
                {errors.description && <p className="text-red-500 text-sm mt-1">{errors.description}</p>}
              </div>

              <SystemPromptSection value={systemPrompt} onChange={(e) => setSystemPrompt(e.target.value)} error={errors.systemPrompt} disabled={isSubmitting} />
              <ModelSelectionSection allModels={AVAILABLE_MODELS} selectedModel={selectedModel} setSelectedModel={setSelectedModel} error={errors.model} isSubmitting={isSubmitting} />
            </CardContent>
          </Card>

          <Card className="bg-[#2A2A2A] border-[#3A3A3A]">
            <CardHeader>
              <CardTitle className="flex items-center gap-2 text-white">
                <FunctionSquare className="h-5 w-5 text-yellow-500" />
                Tools
              </CardTitle>
            </CardHeader>
            <CardContent>
              <ToolsSection allTools={allTools} selectedTools={selectedTools} setSelectedTools={setSelectedTools} isSubmitting={isSubmitting}></ToolsSection>
            </CardContent>
          </Card>

          <div className="flex justify-end">
            <Button className="bg-violet-500 hover:bg-violet-600 text-white" onClick={handleCreateAgent} disabled={isSubmitting}>
              {isSubmitting ? (
                <>
                  <Loader2 className="h-4 w-4 mr-2 animate-spin" />
                  Creating...
                </>
              ) : (
                "Create Agent"
              )}
            </Button>
          </div>
        </div>
      </div>
    </div>
  );
}
