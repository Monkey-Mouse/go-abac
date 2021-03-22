# rules

for better extensible, the type of rules in the model(`RulesType`) is designed as array of interface   
we suggest user implementing handler method `JudgeRule()` to check if the request of subject satisfy the rule  
and hereby, this package provides method `Can()` to return whether the request meet the rules above


## JudgeRule()(bool,error)



## Can() 
### logic

**!!! or**

check each item in the rules array, if **any** rule satisfy(`JudgeRule()`return`true,nil`), 
the request have right to access the resources