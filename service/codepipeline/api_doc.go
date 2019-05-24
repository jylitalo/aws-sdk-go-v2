// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package codepipeline provides the client and types for making API
// requests to CodePipeline.
//
// Overview
//
// This is the AWS CodePipeline API Reference. This guide provides descriptions
// of the actions and data types for AWS CodePipeline. Some functionality for
// your pipeline is only configurable through the API. For additional information,
// see the AWS CodePipeline User Guide (https://docs.aws.amazon.com/codepipeline/latest/userguide/welcome.html).
//
// You can use the AWS CodePipeline API to work with pipelines, stages, actions,
// and transitions, as described below.
//
// Pipelines are models of automated release processes. Each pipeline is uniquely
// named, and consists of stages, actions, and transitions.
//
// You can work with pipelines by calling:
//
//    * CreatePipeline, which creates a uniquely-named pipeline.
//
//    * DeletePipeline, which deletes the specified pipeline.
//
//    * GetPipeline, which returns information about the pipeline structure
//    and pipeline metadata, including the pipeline Amazon Resource Name (ARN).
//
//    * GetPipelineExecution, which returns information about a specific execution
//    of a pipeline.
//
//    * GetPipelineState, which returns information about the current state
//    of the stages and actions of a pipeline.
//
//    * ListPipelines, which gets a summary of all of the pipelines associated
//    with your account.
//
//    * ListPipelineExecutions, which gets a summary of the most recent executions
//    for a pipeline.
//
//    * StartPipelineExecution, which runs the the most recent revision of an
//    artifact through the pipeline.
//
//    * UpdatePipeline, which updates a pipeline with edits or changes to the
//    structure of the pipeline.
//
// Pipelines include stages. Each stage contains one or more actions that must
// complete before the next stage begins. A stage will result in success or
// failure. If a stage fails, then the pipeline stops at that stage and will
// remain stopped until either a new version of an artifact appears in the source
// location, or a user takes action to re-run the most recent artifact through
// the pipeline. You can call GetPipelineState, which displays the status of
// a pipeline, including the status of stages in the pipeline, or GetPipeline,
// which returns the entire structure of the pipeline, including the stages
// of that pipeline. For more information about the structure of stages and
// actions, also refer to the AWS CodePipeline Pipeline Structure Reference
// (https://docs.aws.amazon.com/codepipeline/latest/userguide/pipeline-structure.html).
//
// Pipeline stages include actions, which are categorized into categories such
// as source or build actions performed within a stage of a pipeline. For example,
// you can use a source action to import artifacts into a pipeline from a source
// such as Amazon S3. Like stages, you do not work with actions directly in
// most cases, but you do define and interact with actions when working with
// pipeline operations such as CreatePipeline and GetPipelineState. Valid action
// categories are:
//
//    * Source
//
//    * Build
//
//    * Test
//
//    * Deploy
//
//    * Approval
//
//    * Invoke
//
// Pipelines also include transitions, which allow the transition of artifacts
// from one stage to the next in a pipeline after the actions in one stage complete.
//
// You can work with transitions by calling:
//
//    * DisableStageTransition, which prevents artifacts from transitioning
//    to the next stage in a pipeline.
//
//    * EnableStageTransition, which enables transition of artifacts between
//    stages in a pipeline.
//
// Using the API to integrate with AWS CodePipeline
//
// For third-party integrators or developers who want to create their own integrations
// with AWS CodePipeline, the expected sequence varies from the standard API
// user. In order to integrate with AWS CodePipeline, developers will need to
// work with the following items:
//
// Jobs, which are instances of an action. For example, a job for a source action
// might import a revision of an artifact from a source.
//
// You can work with jobs by calling:
//
//    * AcknowledgeJob, which confirms whether a job worker has received the
//    specified job,
//
//    * GetJobDetails, which returns the details of a job,
//
//    * PollForJobs, which determines whether there are any jobs to act upon,
//
//    * PutJobFailureResult, which provides details of a job failure, and
//
//    * PutJobSuccessResult, which provides details of a job success.
//
// Third party jobs, which are instances of an action created by a partner action
// and integrated into AWS CodePipeline. Partner actions are created by members
// of the AWS Partner Network.
//
// You can work with third party jobs by calling:
//
//    * AcknowledgeThirdPartyJob, which confirms whether a job worker has received
//    the specified job,
//
//    * GetThirdPartyJobDetails, which requests the details of a job for a partner
//    action,
//
//    * PollForThirdPartyJobs, which determines whether there are any jobs to
//    act upon,
//
//    * PutThirdPartyJobFailureResult, which provides details of a job failure,
//    and
//
//    * PutThirdPartyJobSuccessResult, which provides details of a job success.
//
// See https://docs.aws.amazon.com/goto/WebAPI/codepipeline-2015-07-09 for more information on this service.
//
// See codepipeline package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codepipeline/
//
// Using the Client
//
// To use CodePipeline with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the CodePipeline client for more information on
// creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/codepipeline/#New
package codepipeline
