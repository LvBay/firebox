import {createRouter, createWebHashHistory} from 'vue-router'
import Myself from '../pages/myself.vue'
import BeforeGame from '../pages/beforeGame.vue'
import DuringGame from '../pages/duringGame.vue'
import AfterGame from '../pages/afterGame.vue'


const routes = [
    {
        path: '/',
        redirect: '/myself'
    },
    {
        path: '/myself',
        name: 'myself',
        component: Myself,
    },
    {
        path: '/beforeGame',
        name: 'beforeGame',
        component: BeforeGame,
    },
    {
        path: '/duringGame',
        name: 'duringGame',
        component: DuringGame,
    },
    {
        path: '/afterGame',
        name: 'afterGame',
        component: AfterGame,
    },
]

const router = createRouter({
    history: createWebHashHistory(),
    routes
})

export default router