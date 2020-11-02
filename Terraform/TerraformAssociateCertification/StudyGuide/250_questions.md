These were taken from here: https://medium.com/bb-tutorials-and-thoughts/250-practice-questions-for-terraform-associate-certification-7a3ccebe6a1a  
Thank you Bhargav!  

- [Understand infrastructure as code (IaC) concepts](#understand-infrastructure-as-code--iac--concepts)
- [Understand Terraform’s purpose (vs other IaC)](#understand-terraform-s-purpose--vs-other-iac-)
- [Understand Terraform basics](#understand-terraform-basics)
- [Use the Terraform CLI (outside of core workflow)](#use-the-terraform-cli--outside-of-core-workflow-)
- [Interact with Terraform modules](#interact-with-terraform-modules)
- [Navigate Terraform workflow](#navigate-terraform-workflow)
- [Implement and maintain state](#implement-and-maintain-state)
- [Read, generate, and modify the configuration](#read--generate--and-modify-the-configuration)
- [Understand Terraform Cloud and Enterprise capabilities](#understand-terraform-cloud-and-enterprise-capabilities)

### Understand infrastructure as code (IaC) concepts
1.
2.
3.
4.
5.
6.
7.
8.
9.

### Understand Terraform’s purpose (vs other IaC)


10. What is multi-cloud deployment?
11. How multi-cloud deployment is useful?
12. What is cloud-agnostic in terms of provisioning tools?
13.
14.
15.
16.
17.

### Understand Terraform basics

18.

19.

20. Where do you put terraform configurations so that you can configure some behaviors of Terraform itself?
- Inside the terraform block
- Only constants can be used inside the terraform block (no variables)

21. Only constants are allowed inside the terraform block. Is this correct?
- Yes

22. What are the Providers?
- A provider is reponsible for understanding API interactions and exposing resources. Most providers configure specific infrastructure on a cloud platform

23. How do you configure a Provider?
- You add it under `required providers` on the terraform config and configure instances of this provider on the `provider` block

24. What are the meta-arguments that are defined by Terraform itself and available for all provider blocks?
- `version` and `alias`

25. What is Provider initialization and why do we need it?
- Initialization downloads and install the provider plugin so that it can be executed with that provider

26. How do you initialize any Provider?
- `terraform init`

27. When you run terraform init command, all the providers are installed in the current working directory. Is this true?
- Yes. Only for the current dir -- other terraform folders can have their on provider settings

28. How do you constrain the provider version?
- By assigning a version to the provider inside the `required_providers` on the terraform config:
```hcl-terraform
terraform {
  required_providers {
    aws = "~> 1.0"
  }
}
```

29. How do you upgrade to the latest acceptable version of the provider?
- `terraform init -upgrade`

30. How many ways you can configure provider versions?
    - Inside the terraform block
    - Inside the provider block with `version`


31. How do you configure Multiple Provider Instances?
- By having multiple provider blocks with aliases

32. Why do we need Multiple Provider instances?
- For instance when you want to have providers in different regions
- Other examples: Targeting multiple docker hosts, multiple consul hosts

33. How do we define multiple Provider configurations?
- Just write multiple `provider` with the same name but different `alias` values 

34. How do you select alternate providers?
- By specifying the provider inside a resource block

35. What is the location of the user plugins directory?
- The provider plugins distributed by HashiCorp are automatically installed by `terraform init` 
- Third party plugins can be manually installed under `~/.terraform.d/plugins`

36. Third-party plugins should be manually installed. Is that true?
- Yes

37. The command terraform init cannot install third-party plugins? True or false?
- True?

38. What is the naming scheme for provider plugins?
- `terraform-provider-<NAME>_vX.Y.Z`

39. What is the CLI configuration File?
- `.terraformrc`

40. Where is the location of the CLI configuration File?
- On Windows, under `%APPDATA%`, in linux, on your home
- The location of this file can also be changed by setting the `TF_CLI_CONFIG_FILE` variable

41. What is Provider Plugin Cache?
- Terraform allows the use of a local cache directory, which then allows each distinct plugin binary to be downloaded only once and reused by multiple terraform configs

42. How do you enable Provider Plugin Cache?
- Edit your CLI configuration file (by default `.terraformrc` on Linux) and define the directory as something like:
- `plugin_cache_dir = "$HOME/.terraform.d/plugin-cache"`

43. When you are using plugin cache you end up growing cache directory with different versions. Whose responsibility to clean it?
- Terraform will never delete those plugins, so it's your responsability

44. Why do we need to initialize the directory? (in the context of a new terraform directory with configs)
- To download the provider plugin executable to be able to run terraform for the providers configured

45. What is the command to initialize the directory?
- `terraform init`

46. If different teams are working on the same configuration. How do you make files to have consistent formatting?
- Use `terraform fmt` consistently

47. If different teams are working on the same configuration. How do you make files to have syntactically valid and internally consistent?
- Use `terraform validate`

48. What is the command to create infrastructure?
- `terraform apply`

49. What is the command to show the execution plan and not apply?
- `terraform plan`

50. How do you inspect the current state of the infrastructure applied?
- `terraform show`

51. If your state file is too big and you want to list the resources from your state. What is the command?
- `terraform state list`

52. What is plug-in based architecture?
- It defines additional features (like providers) as feature separate from the core of terraform
- This allows the plugin to have a different development schedule

53. What are Provisioners?
- Pretty much a shell script that you can run on your local or a remote machine to execute extra steps after instance creation, destruction or update
- Provisiones are a last resort and their general use is discouraged

54. How do you define provisioners?
- Put a provision block inside a resource block. Multiple provisioner blocks can be added to a given resource

55. What are the types of provisioners?
- local-exec and remote-exec

56. What is a local-exec provisioner and when do we use it?
- It's a provisioner that runs on your local machine

57. What is a remote-exec provisioner and when do we use it?
- It's a provisioner that runs on the remote machine

58. Are provisioners runs only when the resource is created or destroyed?
- Provisioners can be triggered on: create or destroy

59. What do we need to use a remote-exec?
- A connection block inside the provisioner of type ssh or winrm

60. When terraform mark the resources are tainted?
- When resource creation works but a provisioner fails (default behaviour)

61. You applied the infrastructure with terraform apply and you have some tainted resources. You run an execution plan now what happens to those tainted resources?
- They will be destroyed and created again

62. Terraform also does not automatically roll back and destroy the resource during the apply when the failure happens. Why?
- Because it goes against the execution plan, which is to create a resource. Terraform will just mark the resource as tainted

63. How do you manually taint a resource?
- `terraform taint resource.id`

64. Does the taint command modify the infrastructure?
- No

65. By default, provisioners that fail will also cause the Terraform apply itself to fail. Is this true?
- Yes

66. By default, provisioners that fail will also cause the Terraform apply itself to fail. How do you change this?
- By changing the `on_failure` variable to `continue` inside the provision block

67. How do you define destroy provisioner and give an example?
- By specifying the `when` value to be: `when = destroy`

68. How do you apply constraints for the provider versions?
- On the `terraform` config block, on the `required_providers` block (for terraform 0.13)
- With the `version` setting on the standalone provider block (for terraform 0.12)

69. What should you use to set both a lower and upper bound on versions for each provider?
 
70. How do you try experimental features?

71. When does the terraform does not recommend using provisions?
- Whenever you can avoid it, they're meant only as a last resort

72. Expressions in provisioner blocks cannot refer to their parent resource by name. Is this true?
- True. You need to use the `self` object for that!

73. What does this symbol version = “~> 1.0” mean when defining versions?
- Means it'll accept versions bigger than 1.0 but lesser than 2.0

74. Terraform supports both cloud and on-premises infrastructure platforms. Is this true?
- Yes

75. Terraform assumes an empty default configuration for any provider that is not explicitly configured. A provider block can be empty. Is this true?
- Yes

76. How do you configure the required version of Terraform CLI can be used with your configuration?
- Inside the terraform configuration blocc?

77. Terraform CLI versions and provider versions are independent of each other. Is this true?
- Yes

78. You are configuring aws provider and it is always recommended to hard code aws credentials in *.tf files. Is this true?
- False!

79. You are provisioning the infrastructure with the command terraform apply and you noticed one of the resources failed. How do you remove that resource without affecting the whole infrastructure?
- `terraform state rm?`

### Use the Terraform CLI (outside of core workflow)
80. What is command fmt?
- A command that formats terraform files to a common standard on a given directory

81. What is the recommended approach after upgrading terraform?
- The answer given by the author didn't convince me much

83. By default, fmt scans the current directory for configuration files. Is this true?
- Yes

84. You are formatting the configuration files and what is the flag you should use to see the differences?
- `terraform fmt -diff`

85. You are formatting the configuration files and what is the flag you should use to process the subdirectories as well?
- `terraform fmt -recursive`

86. You are formatting configuration files in a lot of directories and you don’t want to see the list of file changes. What is the flag that you should use?
- `terraform fmt --list=false`

87. What is the command taint?
- The command taint marks a resource for destroy and re-creation the next time the apply command runes

88. What is the command usage?
- `terraform taint [options] <resource>`

89. When you are tainting a resource terraform reads the default state file terraform.tfstate. What is the flag you should use to read from a different path?
- `terraform taint -state=path`

90. Give an example of tainting a single resource?
- `terraform taint aws_instance.example1`

91. Give an example of tainting a resource within a module?
- `terraform taint module.MODULENAME.RESOURCENAME`

92. What is the command import?
- It imports a given resource into your terraform state. Note that this resource has to already be defined on a terraform file

93. What is the command import usage?
- `terraform import [options] ADDR ID`

94. What is the default workspace name?
- default

95. What are workspaces?
- Workspaces are like different instances of a state for a given set of terraform configurations

96. What is the command to list the workspaces?
- `terraform workspace list`

97. What is the command to create a new workspace?
- `terraform workspace new banana`

98. What is the command to show the current workspace?
- `terraform workspace show`

99. What is the command to switch the workspace?
- `terraform workspace select banana`

100. What is the command to delete the workspace?
- `terraform workspace delete banana`

101. Can you delete the default workspace?
- No

102. You are working on the different workspaces and you want to use a different number of instances based on the workspace. How do you achieve that?
103. You are working on the different workspaces and you want to use tags based on the workspace. How do you achieve that?
104. You want to create a parallel, distinct copy of a set of infrastructure in order to test a set of changes before modifying the main production infrastructure. How do you achieve that?

105. What is the command state?
- It is the command to manipulate state

106. What is the command usage?
- Eh, see the `cmd_notes.md` file for that

107. You are working on terraform files and you want to list all the resources. What is the command you should use?
108. How do you list the resources for the given name?
109. What is the command that shows the attributes of a single resource in the state file?
110. How do you do debugging terraform?
111. If terraform crashes where should you see the logs?
112. What is the first thing you should do when the terraform crashes?
113. You are building infrastructure for different environments for example test and dev. How do you maintain separate states?
114. What is the difference between directory-separated and workspace-separated environments?
115. What is the command to pull the remote state?
116. What is the command is used manually to upload a local state file to a remote state
117. The command terraform taint modifies the state file and doesn’t modify the infrastructure. Is this true?
118. Your team has decided to use terraform in your company and you have existing infrastructure. How do you migrate your existing resources to terraform and start using it?
119. When you are working with the workspaces how do you access the current workspace in the configuration files?
120. When you are using workspaces where does the Terraform save the state file for the local state?
121. When you are using workspaces where does the Terraform save the state file for the remote state?
122. How do you remove items from the Terraform state?
123. How do you move the state from one source to another?
124. How do you rename a resource in the terraform state file?

### Interact with Terraform modules
125. Where do you find and explore terraform Modules?
- On the terraform public module registry?

126. How do you make sure that modules have stability and compatibility?
- By checking only by verified modules on the terraform registry

127. How do you download any modules?
- Add them on your config under the module block and do `terraform init`

128. What is the syntax for referencing a registry module?
- `<NAMESPACE>/<NAME>/<PROVIDER>`
```hcl-terraform
module "consul" {
  source = "hashicorp/consul/aws"
  version = "0.1.0"
}
```

129. What is the syntax for referencing a private registry module?
- <HOSTNAME>/<NAMESPACE>/<NAME>/<PROVIDER>
```hcl-terraform
module "vpc" {
  source = "app.terraform.io/example_corp/vpc/aws"
  version = "0.9.3"
}
```

130. The terraform recommends that all modules must follow semantic versioning. Is this true?

131. What is a Terraform Module?

132. Why do we use modules for?

133. How do you call modules in your configuration?

134. How many ways you can load modules?

135. What are the best practices for using Modules?

136. What are the different source types for calling modules?

137. What are the arguments you need for using modules in your configuration?

138. How do you set input variables for the modules?

139. How do you access output variables from the modules?

140. Where do you put output variables in the configuration?

141. How do you pass input variables in the configuration?

142. What is the child module?

143. When you use local modules you don’t have to do the command init or get every time there is a change in the local module. why?

144. When you use remote modules what should you do if there is a change in the module?

145. A simple configuration consisting of a single directory with one or more .tf files is a module. Is this true?

146. When using a new module for the first time, you must run either terraform init or terraform get to install the module. Is this true?

147. When installing the modules and where does the terraform save these modules?

148. What is the required argument for the module?

149. What are the other optional meta-arguments along with the source when defining modules

### Navigate Terraform workflow
150. What is the Core Terraform workflow?

151. What is the workflow when you work as an Individual Practitioner?

152. What is the workflow when you work as a team?

153. What is the workflow when you work as a large organization?

154. What is the command init?

155. You recently joined a team and you cloned a terraform configuration files from the version control system. What is the first command you should use?

156. What is the flag you should use to upgrade modules and plugins a part of their respective installation steps?

157. When you are doing initialization with terraform init, you want to skip backend initialization. What should you do?

158. When you are doing initialization with terraform init, you want to skip child module installation. What should you do?

159. When you are doing initialization where do all the plugins stored?
- Inside the `.terraform` directory

160. When you are doing initialization with terraform init, you want to skip plugin installation. What should you do?

161. What does the command terraform validate does?

162. What does the command plan do?

163. What does the command apply do?

164. You are applying the infrastructure with the command apply and you don’t want to do interactive approval. Which flag should you use?

165. What does the command destroy do?

166. How do you preview the behavior of the command terraform destroy?

167. What are implicit and explicit dependencies?

168. Give an example of implicit dependency?

169. Give an example of explicit dependency?

170. How do you save the execution plan?

171. You have started writing terraform configuration and you are using some sample configuration as a basis. How do you copy the example configuration into your working directory?

172. What is the flag you should use with the terraform plan to get detailed on the exit codes?

173. How do you target only specific resources when you run a terraform plan?

174. How do you update the state prior to checking differences when you run a terraform plan?

175. The behavior of any terraform destroy command can be previewed at any time with an equivalent terraform plan -destroy command. Is this true?

176. You have the following file and created two resources docker_image and docker_container with the command terraform apply and you go to the terminal and delete the container with the command docker rm. You come back to your configuration and run the command again. Does terraform recreates the resource?

177. You created a VM instance on AWS cloud provider with the terraform configuration and you log in AWS console and removed the instance. What does the next apply do?

178. You have the following file and created two resources docker_image and docker_container with the command terraform planand you go to the terminal and delete the container with the command docker rm. You come back to your configuration and run the command again. What is the output of the command plan?

### Implement and maintain state
179. What are Backends?

180. What is local Backend?

181. What is the default path for the local backend?

182. What is State Locking?

183. Does Terraform continue if state locking fails?

184. Can you disable state locking?

185. What are the types of Backend?

186. What are remote Backends?

187. What is the benefit of using remote backend?

188. If you want to switch from using remote backend to local backend. What should you do?

189. What does the command refresh do?

190. Does the command refresh modify the infrastructure?

191. How do you backup the state to the remote backend?

192. What is a partial configuration in terms of configuring Backends?

193. What are the ways to provide remaining arguments when using partial configuration?

194. What is the basic requirement when using partial configuration?

195. Give an example of passing partial configuration with Command-line Key/Value pairs?

196. How to unconfigure a backend?

197. How do you encrypt sensitive data in the state?

198. Backends are completely optional. Is this true?

199. What are the benefits of Backends?

200. Why should you be very careful with the Force unlocking the state?

201. You should only use force unlock command when automatic unlocking fails. Is this true?

### Read, generate, and modify the configuration

202. How do you define a variable?
- Using the variable block

203. How do you access the variable in the configuration?

204. How many ways you can assign variables in the configuration?

205. Does environment variables support List and map types?

206. How do you provision infrastructure in a staging environment or a production environment using the same Terraform configuration?

207. How do you assign default values to variables?

208. What are the data types for the variables?

209. Give an example of data type List variables?

210. Give an example of data type Map variables?

211. What is the Variable Definition Precedence?

212. What are the output variables?

213. Hoe do you define an output variable?

214. How do you view outputs and queries them?

215. What are the dynamic blocks?

216. What are the best practices for dynamic blocks?

217. What are the Built-in Functions?

218. Does Terraform language support user-defined functions?

219. What is the built-in function to change string to a number?

220. What is the built-in function to evaluates given expression and returns a boolean whether the expression produced a result without any errors?

221. What is the built-in function to evaluates all of its argument expressions in turn and returns the result of the first one that does not produce any errors?

222. What is Resource Address?

223. What is the Module path?

224. What is the Resource spec?

225. What are complex types and what are the collection types Terraform supports?

226. What are the named values available and how do we refer to?

227. What is the built-in function that reads the contents of a file at the given path and returns them as a base64-encoded string?

228. What is the built-in function that converts a timestamp into a different time format?

229. What is the built-in function encodes a given value to a string using JSON syntax?

230. What is the built-in function that calculates a full host IP address for a given host number within a given IP network address prefix?

### Understand Terraform Cloud and Enterprise capabilities
231. What is Sentinel?
- Sentinel is a policy as a code tool that proactively forbids the creation of infrastructure that is outside policy bounds

232. What is the benefit of Sentinel?
- Codifying policy removes the need for ticketing queues, without sacrificing enforcement.
- One of the other benefits of Sentinel is that it also has a full testing framework.
- Avoiding a ticketing workflow allows organizations to provide more self-service capabilities and end-to-end automation, minimizing the friction for developers and operators.

233. What is the Private Module Registry?

234. What is the difference between public and private module registries when defined source?

235. Where is the Terraform Module Registry available at?

236. What is a workspace?

237. What are the benefits of workspaces?

238. You are configuring a remote backend in the terraform cloud. You didn’t create an organization before you do terraform init. Does it work?

239. You are configuring a remote backend in the terraform cloud. You didn’t create a workspace before you do terraform init. Does it work?

240. Terraform workspaces when you are working with CLI and Terraform workspaces in the Terraform cloud. Is this correct?

241. How do you authenticate the CLI with the terraform cloud?

242. You are building infrastructure on your local machine and you changed your backend to remote backend with the Terraform cloud. What should you do to migrate the state to the remote backend?

243. How do you configure remote backend with the terraform cloud?

244. What is Run Triggers?

245. What is the benefit of Run Triggers?

246. What are the available permissions that terraform clouds can have?

247. Who can grant permissions on the workspaces?

248. Which plan do you need to manage teams on Terraform cloud?

249. How can you add users to an organization?

250. The Terraform Cloud Team plan charges you on a per-user basis. Is this true?
