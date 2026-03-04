# Lists

Response Types:

- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#List">List</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#Pagination">Pagination</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListGetResponse">ListGetResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListListResponse">ListListResponse</a>

Methods:

- <code title="get /lists/{listId}">client.Lists.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, listID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListGetResponse">ListGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /lists">client.Lists.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListListParams">ListListParams</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListListResponse">ListListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Subscribers

Response Types:

- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#Subscriber">Subscriber</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListSubscriberListResponse">ListSubscriberListResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListSubscriberAddResponse">ListSubscriberAddResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListSubscriberRemoveResponse">ListSubscriberRemoveResponse</a>

Methods:

- <code title="get /lists/{listId}/subscribers">client.Lists.Subscribers.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListSubscriberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, listID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListSubscriberListResponse">ListSubscriberListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /lists/{listId}/subscribers">client.Lists.Subscribers.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListSubscriberService.Add">Add</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, listID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListSubscriberAddParams">ListSubscriberAddParams</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListSubscriberAddResponse">ListSubscriberAddResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /lists/{listId}/subscribers/{subscriberId}">client.Lists.Subscribers.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListSubscriberService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subscriberID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListSubscriberRemoveParams">ListSubscriberRemoveParams</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#ListSubscriberRemoveResponse">ListSubscriberRemoveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Customers

Response Types:

- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#Customer">Customer</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CustomerGetResponse">CustomerGetResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CustomerUpdateResponse">CustomerUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CustomerListResponse">CustomerListResponse</a>

Methods:

- <code title="get /customers/{customerId}">client.Customers.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CustomerService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CustomerGetResponse">CustomerGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /customers/{customerId}">client.Customers.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CustomerService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, customerID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CustomerUpdateParams">CustomerUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CustomerUpdateResponse">CustomerUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customers">client.Customers.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CustomerService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CustomerListParams">CustomerListParams</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CustomerListResponse">CustomerListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Campaigns

Response Types:

- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#Campaign">Campaign</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CampaignGetResponse">CampaignGetResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CampaignListResponse">CampaignListResponse</a>

Methods:

- <code title="get /campaigns/{campaignId}">client.Campaigns.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CampaignService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, campaignID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CampaignGetResponse">CampaignGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /campaigns">client.Campaigns.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CampaignService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CampaignListParams">CampaignListParams</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CampaignListResponse">CampaignListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Forms

Response Types:

- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#Form">Form</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormGetResponse">FormGetResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormUpdateResponse">FormUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormListResponse">FormListResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormDeleteResponse">FormDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormListSubmissionsResponse">FormListSubmissionsResponse</a>

Methods:

- <code title="get /forms/{formId}">client.Forms.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, formID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormGetResponse">FormGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /forms/{formId}">client.Forms.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, formID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormUpdateParams">FormUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormUpdateResponse">FormUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /forms">client.Forms.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormListParams">FormListParams</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormListResponse">FormListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /forms/{formId}">client.Forms.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, formID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormDeleteResponse">FormDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /forms/{formId}/submissions">client.Forms.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormService.ListSubmissions">ListSubmissions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, formID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormListSubmissionsParams">FormListSubmissionsParams</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#FormListSubmissionsResponse">FormListSubmissionsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# KnowledgeBases

Response Types:

- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBase">KnowledgeBase</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseGetResponse">KnowledgeBaseGetResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseUpdateResponse">KnowledgeBaseUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseListResponse">KnowledgeBaseListResponse</a>

Methods:

- <code title="get /knowledge-bases/{kbId}">client.KnowledgeBases.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, kbID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseGetResponse">KnowledgeBaseGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /knowledge-bases/{kbId}">client.KnowledgeBases.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, kbID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseUpdateParams">KnowledgeBaseUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseUpdateResponse">KnowledgeBaseUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /knowledge-bases">client.KnowledgeBases.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseListResponse">KnowledgeBaseListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /knowledge-bases/{kbId}">client.KnowledgeBases.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, kbID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

## Articles

Response Types:

- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseArticle">KnowledgeBaseArticle</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseArticleNewResponse">KnowledgeBaseArticleNewResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseArticleListResponse">KnowledgeBaseArticleListResponse</a>

Methods:

- <code title="post /knowledge-bases/{kbId}/articles">client.KnowledgeBases.Articles.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseArticleService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, kbID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseArticleNewParams">KnowledgeBaseArticleNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseArticleNewResponse">KnowledgeBaseArticleNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /knowledge-bases/{kbId}/articles">client.KnowledgeBases.Articles.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseArticleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, kbID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseArticleListParams">KnowledgeBaseArticleListParams</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#KnowledgeBaseArticleListResponse">KnowledgeBaseArticleListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Pages

Response Types:

- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#Page">Page</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#PageGetResponse">PageGetResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#PageUpdateResponse">PageUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#PageListResponse">PageListResponse</a>
- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#PageDeleteResponse">PageDeleteResponse</a>

Methods:

- <code title="get /pages/{pageId}">client.Pages.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#PageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pageID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#PageGetResponse">PageGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /pages/{pageId}">client.Pages.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#PageService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pageID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#PageUpdateParams">PageUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#PageUpdateResponse">PageUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /pages">client.Pages.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#PageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#PageListParams">PageListParams</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#PageListResponse">PageListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /pages/{pageId}">client.Pages.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#PageService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, pageID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#PageDeleteResponse">PageDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
