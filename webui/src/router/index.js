import {createRouter, createWebHashHistory} from 'vue-router'

const router = createRouter({
    history: createWebHashHistory(import.meta.env.BASE_URL),
    routes: [
        {path: '/', name: 'LoginView', component: () => import('../views/LoginView.vue')},
        {path: '/logged', name: 'SuccesView', component: () => import('../views/SuccessView.vue')},
        {path: '/error', name: 'ErrorView', component: () => import('../views/ErrorView.vue')},
    ]

})

export default router