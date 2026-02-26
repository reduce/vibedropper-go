# Lists

Response Types:

- <a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#List">List</a>
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
- <code title="get /campaigns">client.Campaigns.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CampaignService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go">vibedropper</a>.<a href="https://pkg.go.dev/github.com/reduce/vibedropper-go#CampaignListResponse">CampaignListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
