AccessControl~IAccessInfo:Object  inner
> An interface that defines an access information to be granted or denied.
> When you start a method chain with `AccessControl#grant`
> or AccessControl#deny methods, you're actually building this object
> which will eventually be committed to the underlying grants model.
#### properties
- action
- possession
- resource
- subject
- rules

AccessControl~IQueryInfo:Object  inner
> An interface that defines an access information to be queried.
> When you start a method chain with `AccessControl#Can` method,
> you're actually building this query object which will be used to
> check the access permissions.
>
>
#### properties
- action
- possession
- resource
- subject