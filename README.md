# reacting-auth

This project showcases a highly efficient Role-Based Access Control (RBAC) system, built with the latest architecture and tools including Gin, Casbin, JWT, microservices, Kubernetes, Etcd, NATS, and microfrontends with React.

The RBAC system ensures secure and scalable management of user roles and permissions. With Gin's powerful routing capabilities and Casbin's flexible access control library, the project delivers efficient authorization management.

Using microservices architecture enables modular development and independent deployment of components. Kubernetes ensures reliable container orchestration, while Etcd serves as a distributed key-value store for configuration management. NATS enables fast and efficient communication between microservices.

The incorporation of microfrontends with React allows for flexible and seamless frontend development, enhancing user experiences.

In summary, this project demonstrates a robust and modern approach to RBAC, leveraging top-notch architecture and tools for efficient access control and scalability.


<p align="center">
    <img src="https://gitee.com/bullteam/zeus-admin/raw/master/docs/images/logo.png" height="145">
</p>

# Zeus - Permission & Account Management System

[![golang](https://img.shields.io/badge/golang-1.18.2-green.svg?style=plastic)](https://www.golang.org/)
[![casbin](https://img.shields.io/badge/casbin-2.47.1-brightgreen.svg?style=plastic)](https://github.com/casbin/casbin)

## Project Introduction
> `Zeus` is a permission management backend system that provides private SaaS cloud services for unified backend permission management for enterprises.  
> - The project is developed using the `Golang Gin` + `vue-element-admin` framework, using `JWT` + `Casbin` for permission management, and provides a Restful API interface for OAuth2.0.
> - It provides unified login authentication, menu management, permission management, organizational structure management, employee management, configuration center, and log management for enterprise backend systems.
> - Supports enterprise WeChat and DingTalk login and synchronization of enterprise organizational structure.
> - Centralizes employee onboarding and offboarding, and strengthens permission approval processes.
> - Integrates open-source software, paid SaaS software, and internal development systems of enterprises, including but not limited to Jenkins, Jira, GitLab, Confluence, ZenTao, enterprise email, OA, CRM, financial software, and enterprise SaaS cloud services, to solve the problem of unsynchronized accounts across multiple software and platforms in enterprises.
> - `Build a unified open platform ecosystem to make it easier for enterprises to introduce external systems.`

For more information, please visit the official website of [Bull Team Open Source](http://www.bullteam.cn) and the detailed [Development Documentation Guide](http://doc.bullteam.cn).

## Features (Currently Implemented)
- Login/Logout
- Permission Management
    - User Management (Personnel Management)
    - Role Management (Function and Permission Management)
    - Department Management
    - Project Management
    - Menu Management
    - Data Permission Management
- Personal Account
    - Third-party Login (DingTalk)
    - Security Settings ([Google 2FA Two-Factor Authentication](http://www.ruanyifeng.com/blog/2017/11/2fa-tutorial.html))
    - LDAP Support

## Roadmap (Planned Features)
- Organizational Structure Management (DingTalk Integration)
- Security and Risk Control
- Operation Log Monitoring
    - Login Logs
    - Abnormal Logins
    - Operation Logs
- Page Management
    - Page Configuration Management
- Configuration Center
- Application Center (Open Platform)
- Personal Account
    - Phone Verification
    - Email Verification
- Add support for enterprise WeChat, WeChat, GitHub, Gmail, QQ, etc., for login
- Login Authorization (OAuth 2.0, Ldap, SAML2.0, Cas, etc.)
- Integration with Worklite, Teambition, Github, Mockplus, Tapd, and other SaaS services
- Integration with open-source software such as Jenkins, Jira, GitLab, Confluence, ZenTao, etc.

# Docker Deployment
You can refer to the [documentation](http://doc.bullteam.cn/guide/install.html#%E4%BB%8Edocker%E5%AE%89%E8%A3%85) for Docker deployment instructions.

# Architecture
<img src="https://
