Go & Svelte basic auth app
===

The purpose of this project is to create a basic auth app using Go and Svelte to skill up on these technologies.

## 1. Backend

### 1.1. Stack

- Go : Language (Compiled language, so it's fast and light + it's really scalable)
- Go Fiber : Backend framework (build on top of fastHTTP, really fast)
- Go Ent : ORM (entity framework Backend - Database)

### 1.2. Init the project
Todo

---


## 2. Front end

### 2.1. Stack

- TypeScript : Language (Typed JavaScript, prevent some error, cleaner code...)
- Svelte : Frontend Framework (Easy to learn, great performances, but quite young)
- Svelte Query 
- Yarn : Package manager
- Vite : Build tool to bundle the code for production
- Tailwind : CSS Library (really sexy and easy to use)
- DaisyUI : CSS Library on top of Tailwind (because I'm really lazy...)
- 

### 2.2. Init the project

```shell
# Init Vite / Svelte.ts project
yarn create vite web # --template svelte

# Go in the new /web/ folder
cd web

# Do the installation
yarn install

# Quick launch & test
yarn dev

# Go on localhost:3000...
# OK, that's sexy !


yarn add ts-node
# Let's add Tailwind & DaisyUI
# https://tailwindcss.com/docs/guides/sveltekit
yarn add tailwindcss postcss autoprefixer daisyui
yarn tailwind init tailwind.config.cjs -p
mv postcss.config.js postcss.config.cjs
# https://daisyui.com/docs/install/
```


---
Thibault **Santonja**<br/>
*2022*