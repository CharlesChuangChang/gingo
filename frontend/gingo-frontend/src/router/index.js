import { createRouter, createWebHashHistory } from "vue-router/dist/vue-router";


const routes = [
    {
        path: '/',
        redirect: '/Home',
    },
    {
        path: '/Login',
        name: "Login",
        component: () => import("../views/LoginPage.vue")
    },
    {
        path: '/Register',
        name: "Register",
        component: () => import("../views/RegisterPage.vue")
    },
    {
        path: '/Home',
        name: "Home",
        component: () => import("../views/HomePage.vue")
    }
];

const router = createRouter({
    mode: 'history',
    history: createWebHashHistory(),
    base: process.env.BASE_URL,
    routes
});

//导航守卫
router.beforeEach((to,from,next)=>{
    if (to.path==='/Login') return next();
    if (to.path==='/Register') return next();
    //获取token
    const tokenStr= window.sessionStorage.getItem('token')
    if(!tokenStr) return next('/Login')
    next()
   
  })
 
export default router;