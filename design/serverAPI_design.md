# serverAPI_Design

> The serverAPI design adopts the RESTful interface design pattern and designs the corresponding API interface according to the specific requirements. 

**From the perspective of requirements, related apis can be divided into the following points:**

## I. inquiry related:


*  **Query nameserver information:**

 check out an agent information: `  GET bootcamp/agent/nameserver? AgentId = XX`

 check out all the agent information: `GET bootcamp/allAgents/nameserver`
 
 

*  **Query historical operation information:**

 the query an agent: `GET bootcamp/agent/historyMessage? AgentId = XX`

 check out all the agent: `GET bootcamp/allAgents/historyMessage`


*  **Query CPU usage:**

 Query an agent: `GET bootcamp/agent/CPU? AgentId = XX`

 Query allAgents:` GET bootcamp/allAgents/CPU`


*  **Query mirror information:**

 the query an agent: `GET bootcamp/agent/imageMessage/agentId = XX`

 check out all the agent: `GET bootcamp/allAgents/imageMessage`


*  **Query permission information:**

 the query of the single agent: `GET bootcamp/agent/permission? AgentId = XX`
 
## Ii. Deletion related:


*  **Delete mirror information:**

Delete an agent image: `Delete bootcamp/agent/image? XX && agentId = imageId = XX`

 Delete one image of all agents: `Delete bootcamp/allAgents/image? ImageId = XX`



## Iii. Modification related:


*  **Change the nameserver information:**

change a agent: `PUT bootcamp/agent/nameserver? AgentId = XX`

change all agent: `PUT bootcamp/allAgents/nameserver`


*  **Change of local domain name resolution information:**

Change an agent: `PUT bootcamp/agent/DNS? AgentId = XX`

Change allAgents: `PUT bootcamp/allAgents/DNS`


* **Download image:**

 Add images to an agent: `PUT bootcamp/agent/image? AgentId = XX`

 Add mirrors to allAgents: `PUT bootcamp/allAgents/image`


* **Rollback history operation: a rollback operation can also be understood as an operation to update backwards**

 Rollback single agent: `PUT bootcamp/agent/rollback? AgentId = XX`

Rollback all the operation of the agent:` PUT bootcamp/allAgents/rollback`



