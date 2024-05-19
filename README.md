# About

- A simple Go program that reads a file and prints the data.
- The data is read from [event.json](event.json) file.
- The program will only print the following fields instead of the whole JSON object:
    - `Records[0].Sns.Message.Region`
    - `Records[0].Sns.Message.Detail`
    - `Records[0].Sns.Message.Detail.Pipeline`
    - `Records[0].Sns.Message.Detail.ExecutionId`
    - `Records[0].Sns.Message.Detail.State`
- This is a Go refresher for me. It's been a while since I last wrote code using this language. 😅

# Usage

```bash
go run main.go
```

# Output

```bash
→ go run main.go
2024/05/20 02:02:19 "event.json" file opened successfully!
2024/05/20 02:02:19 Region: us-west-1
2024/05/20 02:02:19 Pipeline Name: dev.example.com-frontend
2024/05/20 02:02:19 Execution ID: 5a7c44a9-8776-448d-b753-00c82a322f6a
2024/05/20 02:02:19 State: SUCCEEDED
2024/05/20 02:02:19 URL: https://us-west-1.console.aws.amazon.com/codesuite/codepipeline/pipelines/dev.example.com-frontend/executions/5a7c44a9-8776-448d-b753-00c82a322f6a/visualization?region=us-west-1
2024/05/20 02:02:19 Message for Slack or Google Chat: dev.example.com-frontend has succeeded. Please check this link for more details: https://us-west-1.console.aws.amazon.com/codesuite/codepipeline/pipelines/dev.example.com-frontend/executions/5a7c44a9-8776-448d-b753-00c82a322f6a/visualization?region=us-west-1
```
