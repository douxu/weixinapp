# API接口文档
## 1.验证登录信息
### POST /api/login/VerifyLoginInfo
#### 期望输入值:
````
{
  userName string  
  userPassword string
}
````
#### 标准化输出项:
````
{
  errCode string  
  errMsg  string  
  token string  
}
````
#### 正确输出项示例:
````
{
  0  
  ""  
  "cf e4 b6 6b d0 d3 4e ce de 23 20 a2 96 d1 da 77 76 e5 c0 55"  
}
````
#### 错误输出项示例:
````
{
  2  
  "Verify logininfo failed"  
  ""  
}
````
## 2.插入或更新基本配置
### POST /api/configure/PostBasicConfig/:userID
#### 期望输入项:
````
{
  token string  
  mainPageTitle string  
  categoryName []string  
  mode string  
  themeColor string  
}
````
#### 标准化输出项:
````
{
  errCode string  
  errMsg  string  
  result  bool  
}
````
#### 正确输出项示例:
````
{
  0  
  ""  
  true  
}
````
#### 错误输出项示例:
````
{
  2  
  "Insert into bd failed"  
  false  
}
````
## 3.获取浏览记录
### GET /api/record/GetBrowseRecord/:articleID
#### 期望输入项:
````
{
  token string
}
````
#### 标准化输出项:
````
{  
  errCode int  
  errMsg string  
  data   []BrowseRecord  
}
type BrowseRecord struct{
  browseName string  
  browseTime string  
}
````
#### 正确输出项示例:
````
{
  0  
  ""  
  {  
    "test"  
    "2017.4.4 18:00"  
  },  
  {  
    "test"  
    "2017.4.5 17:00"  
  }  
}
````
#### 错误输出项示例:
````
{
  2  
  "Not exist thsi articleID in db"  
  ""  
}
````
## 4.获取留资信息
### GET /api/record/GetFillInfoRecord/:articleID
#### 期望输入项:
````
{
  token string
}
````
#### 标准化输出项:
````
{
  errCode int  
  errMsg string  
  data   []string  
}
type FillInfoRecord struct{  
  informationlist string  
}
````
#### 正确输出项示例:
````
{
  0  
  ""  
  {  
    "收入调查","test,27，2000","test1,30，10000"  
  }  
}  
````
#### 错误输出项示例:
````
{  
  2  
  "No exist this articleID in db"  
  ""  
}  
````
## 5.获取文章列表
### GET /api/record/GetArticleList/:userID
#### 期望输入项:
````
{
  token string  
}
````
#### 标准化输出项:
````
{
  errCode int  
  errMsg string  
  data   []ArticleRecord  
}
type ArticleRecord struct{  
  articleTitle string  
  articleURL string  
  articleStatus bool  
}
````
#### 正确输出项示例:
````
{
  0  
  ""  
  {  
    "test"  
    "xxxxx"  
    "true"  
  },  
  {  
    "test"  
    "xxxxx"  
    "false"  
  }  
}
````
#### 错误输出项示例:
````
{  
  2  
  "No exist this userID in db"  
  ""  
}
````
## 6.获取文章详情
### GET /api/content/GetArticleInfo/:articleID
#### 期望输入值:
````
{
  token string  
}
````
#### 标准化输出项:
````
{
  errCode int  
  errMsg string  
  data   ArticleDetailInfo  
}  
  type ArticleDetailInfo struct {
  userID                 int  
  mainPageTitle          string  //文章主题  
  pictureURL             string  //文章图片URL  
  articleContent         string  //文章正文  
  readTargetNum          int     //文章目标阅读数  
  increaseStartTime      int     //增长开始时间  
  increaseContinuousTime int     //增长持续时间  
  recommandOrDataset     string  //推荐或留资设置  
}
````
#### 正确输出项示例:
````
{
  0  
  ""  
  {  
    12345  
    "test"  
    "URL"  
    "正文内容"  
    "1300"  
    "2017.4.3 12:00"  
    "12"  
    "收入调查，姓名，性别，年龄，收入"  
  }  
}
````
#### 错误输出项示例:
````
{
  2  
  "No exist this articleID in db"  
  ""  
}
````
## 7.插入浏览记录
### POST /api/record/InsertBrowseRecord/:articleID
#### 期望输入项:
````
{  
  token string  
  browseName string  
  browseTime string  
}
````
#### 标准化输出项:
````
{
  errCode int  
  errMsg string  
  data   []BrowseRecord  
}
type BrowseRecord struct {
  browseName string  
  browseTime string  
}
````
#### 正确输出项示例:
````
{
  0  
  ""  
  {  
    {  
    "test"  
    "2017.4.3 12:00"  
    },  
    {  
    "test"  
    "2017.4.3 12:00"  
    }  
  }  
}
````
#### 错误输出项示例:
````
{
  2  
  "No exist this articleID in db"  
  ""  
}
````
## 8.插入留资信息
### POST /api/record/InsertFillInfoRecord/:articleID
#### 期望输入项:
````
{
  token string  
  informationList string  
}
````
#### 标准输出项:
````
{
  errCode int  
  errMsg string  
  data bool  
}
````
#### 正确输出项示例:
````
{
  0  
  ""  
  true  
}
````
#### 错误输出项示例:
````
{
  2  
  "No exist this articleID in db"  
  false  
}
````
## 9.删除文章
### POST /api/content/DeleteArticle/:articleID
#### 期望输入项:
````
{
  token string  
}
````
#### 标准输出项:
````
{
  errCode int  
  errMsg string  
  data bool  
}
````
#### 正确输出项示例:
```
{  
  0  
  ""  
  true  
}
```
#### 错误输出项示例:
````
{  
  2  
  "No exist this articleID in db"  
  false  
}
````
