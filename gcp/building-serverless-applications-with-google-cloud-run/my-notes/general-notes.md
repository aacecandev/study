a policy binds a member to a role, which is a list of permissions. The policy is added to a resource

There are two separate credentials managed by the gcloud CLI:

    When you execute gcloud commands, gcloud uses the credentials that are saved when you ran gcloud auth login.

    When you talk to Google Cloud from a client library in your software or an application like Terraform, it uses the application default credentials. Those are the credentials that were saved when you ran gcloud auth application-default login.
