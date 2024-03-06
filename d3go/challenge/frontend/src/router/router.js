import { createRouter, createWebHashHistory } from "vue-router";

const routes = [{
    path: "/loginPage",
    name: "loginPage",
    component: () =>
        import ("../modules/login/LoginPage.vue")
}, {
    path: "/registerPage",
    name: "registerPage",
    component: () =>
        import ("../modules/register/RegisterPage.vue")
}, {
    path: "/uploadPage",
    name: "uploadPage",
    component: () =>
        import ("../modules/upload/UploadPage.vue")
}, {
    path: "/:pathMatch(.*)*",
    redirect: "/loginPage"
}]

const router = createRouter({
    history: createWebHashHistory(),
    routes,
})

router.beforeEach((to, from, next) => {
    const mysession = document.cookie;
    if (to.path !== '/loginPage' && !mysession.includes('mysession=') && to.path !== '/registerPage') {
        next('/loginPage');
    } else {
        next();
    }
});

export default router;