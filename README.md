# blog
个人博客后端

GET     user/info/:name                     获取用户公开信息  
GET     user/detail/:name                   获取用户全部信息  
GET     article/view/:id                    通过id获取文章  
GET     article/titles?number=&page=&type=  获取指定条数的文章  
GET     article/review?number=&page=        获取待审核文章  
POST    article  
POST    login  
POST    register  
POST    article/review                      审核通过或关闭文章