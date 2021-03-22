## [wiki](https://en.wikipedia.org/wiki/Attribute-based_access_control)

### Attributes
> Attributes can be about anything and anyone. They tend to fall into 4 different categories:

- Subject attributes: attributes that describe the user attempting the access e.g. age, clearance, department, role, job title...
- Action attributes: attributes that describe the action being attempted e.g. read, delete, view, approve...
- Object attributes: attributes that describe the object (or resource) being accessed e.g. the object type (medical record, bank account...), the department, the classification or sensitivity, the location...
- Contextual (environment) attributes: attributes that deal with time, location or dynamic aspects of the access control scenario[7]


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