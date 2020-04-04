# Step Functions

## Basics

- you can start multiple executions of the state machine concurrently.

* execution can be triggered **manually using API**, using **APIGW** or also using **CloudWatch Events**

- step functions are quite costly.

## When to use Step Functions

- when you want to visualize workflows, show them to people who are not engineers

* when you want A LOT of control around errors and retries

- for auditing

### Overall:
  
    - business critical workflows (visibility, error handling)
    - complex workflow with branching logic
    - long-running workflows
