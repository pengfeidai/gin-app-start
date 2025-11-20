# API 接口文档

## 基本信息

- **Base URL**: `http://localhost:9060`
- **API Version**: v1
- **API Prefix**: `/api/v1`
- **Content-Type**: `application/json`

## 通用说明

### 响应格式

所有接口返回统一的 JSON 格式：

```json
{
  "code": 0,
  "message": "success",
  "data": {}
}
```

### 状态码说明

| HTTP Status | Code | Description |
|-------------|------|-------------|
| 200 | 0 | 成功 |
| 200 | 10001 | 参数错误 |
| 200 | 10002 | 用户不存在 |
| 200 | 10003 | 未授权访问 |
| 200 | 10004 | 用户已存在 |
| 200 | 10005 | 数据库错误 |
| 200 | 42900 | 请求过于频繁 |
| 200 | 50000 | 系统内部错误 |

## 接口列表

### 1. 健康检查

#### 1.1 服务健康检查

检查服务是否正常运行。

**接口地址**: `GET /health`

**请求参数**: 无

**响应示例**:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "status": "ok",
    "message": "service is running"
  }
}
```

**cURL 示例**:

```bash
curl http://localhost:9060/health
```

---

### 2. 用户管理

#### 2.1 创建用户

创建新用户账号。

**接口地址**: `POST /api/v1/users`

**请求头**:

```
Content-Type: application/json
```

**请求参数**:

| 字段 | 类型 | 必填 | 说明 | 限制 |
|------|------|------|------|------|
| username | string | 是 | 用户名 | 3-32字符 |
| email | string | 否 | 邮箱地址 | 有效的邮箱格式 |
| phone | string | 否 | 手机号码 | 11位数字 |
| password | string | 是 | 密码 | 6-32字符 |

**请求示例**:

```json
{
  "username": "johndoe",
  "email": "john@example.com",
  "phone": "13800138000",
  "password": "securePass123"
}
```

**响应示例**:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "created_at": "2025-11-20T10:00:00Z",
    "update_at": "2025-11-20T10:00:00Z",
    "username": "johndoe",
    "email": "john@example.com",
    "phone": "13800138000",
    "avatar": "",
    "status": 1
  }
}
```

**错误响应**:

```json
{
  "code": 10004,
  "message": "User already exists",
  "data": null
}
```

**cURL 示例**:

```bash
curl -X POST http://localhost:9060/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "email": "john@example.com",
    "phone": "13800138000",
    "password": "securePass123"
  }'
```

---

#### 2.2 获取用户信息

根据用户ID获取用户详细信息。

**接口地址**: `GET /api/v1/users/:id`

**路径参数**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | integer | 是 | 用户ID |

**请求参数**: 无

**响应示例**:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "created_at": "2025-11-20T10:00:00Z",
    "update_at": "2025-11-20T10:00:00Z",
    "username": "johndoe",
    "email": "john@example.com",
    "phone": "13800138000",
    "avatar": "http://example.com/avatar.jpg",
    "status": 1
  }
}
```

**错误响应**:

```json
{
  "code": 10002,
  "message": "User not found",
  "data": null
}
```

**cURL 示例**:

```bash
curl http://localhost:9060/api/v1/users/1
```

---

#### 2.3 更新用户信息

更新指定用户的信息。

**接口地址**: `PUT /api/v1/users/:id`

**路径参数**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | integer | 是 | 用户ID |

**请求头**:

```
Content-Type: application/json
```

**请求参数**:

| 字段 | 类型 | 必填 | 说明 | 限制 |
|------|------|------|------|------|
| email | string | 否 | 邮箱地址 | 有效的邮箱格式 |
| phone | string | 否 | 手机号码 | 11位数字 |
| avatar | string | 否 | 头像URL | 有效的URL格式 |
| status | integer | 否 | 用户状态 | 0=禁用, 1=正常 |

**请求示例**:

```json
{
  "email": "newemail@example.com",
  "phone": "13900139000",
  "avatar": "http://example.com/new-avatar.jpg",
  "status": 1
}
```

**响应示例**:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "id": 1,
    "created_at": "2025-11-20T10:00:00Z",
    "update_at": "2025-11-20T11:00:00Z",
    "username": "johndoe",
    "email": "newemail@example.com",
    "phone": "13900139000",
    "avatar": "http://example.com/new-avatar.jpg",
    "status": 1
  }
}
```

**cURL 示例**:

```bash
curl -X PUT http://localhost:9060/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newemail@example.com",
    "phone": "13900139000"
  }'
```

---

#### 2.4 删除用户

删除指定用户（软删除）。

**接口地址**: `DELETE /api/v1/users/:id`

**路径参数**:

| 参数 | 类型 | 必填 | 说明 |
|------|------|------|------|
| id | integer | 是 | 用户ID |

**请求参数**: 无

**响应示例**:

```json
{
  "code": 0,
  "message": "Deleted successfully",
  "data": null
}
```

**cURL 示例**:

```bash
curl -X DELETE http://localhost:9060/api/v1/users/1
```

---

#### 2.5 用户列表

获取用户列表，支持分页。

**接口地址**: `GET /api/v1/users`

**查询参数**:

| 参数 | 类型 | 必填 | 说明 | 默认值 |
|------|------|------|------|--------|
| page | integer | 否 | 页码 | 1 |
| page_size | integer | 否 | 每页数量 | 10 |

**请求参数**: 无

**响应示例**:

```json
{
  "code": 0,
  "message": "success",
  "data": {
    "list": [
      {
        "id": 1,
        "created_at": "2025-11-20T10:00:00Z",
        "update_at": "2025-11-20T10:00:00Z",
        "username": "johndoe",
        "email": "john@example.com",
        "phone": "13800138000",
        "avatar": "",
        "status": 1
      },
      {
        "id": 2,
        "created_at": "2025-11-20T10:05:00Z",
        "update_at": "2025-11-20T10:05:00Z",
        "username": "janedoe",
        "email": "jane@example.com",
        "phone": "13900139000",
        "avatar": "",
        "status": 1
      }
    ],
    "total": 50,
    "page": 1,
    "page_size": 10
  }
}
```

**cURL 示例**:

```bash
# 获取第1页，每页10条
curl "http://localhost:9060/api/v1/users?page=1&page_size=10"

# 获取第2页，每页20条
curl "http://localhost:9060/api/v1/users?page=2&page_size=20"
```

---

## Postman 集合

### 导入 Postman

将以下 JSON 保存为 `gin-app-start.postman_collection.json` 并导入到 Postman：

```json
{
  "info": {
    "name": "Gin App Start API",
    "schema": "https://schema.getpostman.com/json/collection/v2.1.0/collection.json"
  },
  "item": [
    {
      "name": "Health Check",
      "request": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "{{baseUrl}}/health",
          "host": ["{{baseUrl}}"],
          "path": ["health"]
        }
      }
    },
    {
      "name": "Create User",
      "request": {
        "method": "POST",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n  \"username\": \"testuser\",\n  \"email\": \"test@example.com\",\n  \"phone\": \"13800138000\",\n  \"password\": \"password123\"\n}"
        },
        "url": {
          "raw": "{{baseUrl}}/api/v1/users",
          "host": ["{{baseUrl}}"],
          "path": ["api", "v1", "users"]
        }
      }
    },
    {
      "name": "Get User",
      "request": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "{{baseUrl}}/api/v1/users/1",
          "host": ["{{baseUrl}}"],
          "path": ["api", "v1", "users", "1"]
        }
      }
    },
    {
      "name": "Update User",
      "request": {
        "method": "PUT",
        "header": [
          {
            "key": "Content-Type",
            "value": "application/json"
          }
        ],
        "body": {
          "mode": "raw",
          "raw": "{\n  \"email\": \"newemail@example.com\",\n  \"phone\": \"13900139000\"\n}"
        },
        "url": {
          "raw": "{{baseUrl}}/api/v1/users/1",
          "host": ["{{baseUrl}}"],
          "path": ["api", "v1", "users", "1"]
        }
      }
    },
    {
      "name": "Delete User",
      "request": {
        "method": "DELETE",
        "header": [],
        "url": {
          "raw": "{{baseUrl}}/api/v1/users/1",
          "host": ["{{baseUrl}}"],
          "path": ["api", "v1", "users", "1"]
        }
      }
    },
    {
      "name": "List Users",
      "request": {
        "method": "GET",
        "header": [],
        "url": {
          "raw": "{{baseUrl}}/api/v1/users?page=1&page_size=10",
          "host": ["{{baseUrl}}"],
          "path": ["api", "v1", "users"],
          "query": [
            {
              "key": "page",
              "value": "1"
            },
            {
              "key": "page_size",
              "value": "10"
            }
          ]
        }
      }
    }
  ],
  "variable": [
    {
      "key": "baseUrl",
      "value": "http://localhost:9060"
    }
  ]
}
```

---

## 错误处理

### 错误码列表

| Code | Message | Description | Solution |
|------|---------|-------------|----------|
| 0 | success | 操作成功 | - |
| 10001 | Invalid parameters | 请求参数验证失败 | 检查请求参数格式和内容 |
| 10002 | User not found | 用户不存在 | 确认用户ID是否正确 |
| 10003 | Unauthorized access | 未授权访问 | 检查认证信息 |
| 10004 | User already exists | 用户已存在 | 更换用户名或邮箱 |
| 10005 | Database error | 数据库操作失败 | 联系管理员 |
| 10010 | Failed to query user | 查询用户失败 | 重试或联系管理员 |
| 10011 | Failed to query email | 查询邮箱失败 | 重试或联系管理员 |
| 10012 | Email already exists | 邮箱已存在 | 更换邮箱地址 |
| 10013 | Failed to create user | 创建用户失败 | 检查参数或联系管理员 |
| 10014 | Failed to get user | 获取用户失败 | 重试或联系管理员 |
| 10015 | Failed to get user | 获取用户失败 | 重试或联系管理员 |
| 10016 | Failed to get user | 获取用户失败 | 重试或联系管理员 |
| 10017 | Failed to update user | 更新用户失败 | 检查参数或联系管理员 |
| 10018 | Failed to delete user | 删除用户失败 | 重试或联系管理员 |
| 10019 | Failed to get user list | 获取用户列表失败 | 重试或联系管理员 |
| 42900 | Too many requests | 请求过于频繁 | 降低请求频率 |
| 50000 | Internal server error | 服务器内部错误 | 联系管理员 |

### 错误处理示例

#### JavaScript/Fetch

```javascript
async function createUser(userData) {
  try {
    const response = await fetch('http://localhost:9060/api/v1/users', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(userData)
    });

    const result = await response.json();

    if (result.code === 0) {
      console.log('User created:', result.data);
      return result.data;
    } else {
      console.error('Error:', result.message);
      throw new Error(result.message);
    }
  } catch (error) {
    console.error('Request failed:', error);
    throw error;
  }
}

// 使用示例
createUser({
  username: 'johndoe',
  email: 'john@example.com',
  password: 'password123'
})
  .then(user => console.log('Success:', user))
  .catch(error => console.error('Failed:', error));
```

#### Python/Requests

```python
import requests

def create_user(user_data):
    url = 'http://localhost:9060/api/v1/users'
    headers = {'Content-Type': 'application/json'}
    
    response = requests.post(url, json=user_data, headers=headers)
    result = response.json()
    
    if result['code'] == 0:
        print('User created:', result['data'])
        return result['data']
    else:
        raise Exception(f"Error {result['code']}: {result['message']}")

# 使用示例
try:
    user = create_user({
        'username': 'johndoe',
        'email': 'john@example.com',
        'password': 'password123'
    })
    print('Success:', user)
except Exception as e:
    print('Failed:', e)
```

---

## 速率限制

- **默认限制**: 100 请求/秒/IP
- **限制类型**: 基于 IP 地址的令牌桶算法
- **超限响应**: HTTP 200, Code 42900
- **配置位置**: `configs/*.yaml` 中的 `server.limit_num`

---

## 数据模型

### User (用户)

| 字段 | 类型 | 说明 | 约束 |
|------|------|------|------|
| id | uint | 用户ID | 主键，自增 |
| created_at | timestamp | 创建时间 | 自动生成 |
| update_at | timestamp | 更新时间 | 自动更新 |
| deleted_at | timestamp | 删除时间 | 软删除标记 |
| username | string | 用户名 | 唯一，3-64字符 |
| email | string | 邮箱 | 唯一，最多128字符 |
| phone | string | 手机号 | 唯一，最多32字符 |
| password | string | 密码（加密） | 不返回 |
| salt | string | 密码盐值 | 不返回 |
| avatar | string | 头像URL | 最多256字符 |
| status | int8 | 状态 | 1=正常, 0=禁用 |

---

## 更新日志

### v2.0.0 (2025-11-20)

- ✅ 重构项目架构
- ✅ 升级到 Go 1.24
- ✅ 支持 PostgreSQL
- ✅ 实现用户管理 CRUD API
- ✅ 添加健康检查接口
- ✅ 实现限流功能

---

**文档版本**: v2.0.0  
**最后更新**: 2025-11-20

