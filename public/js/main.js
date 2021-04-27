const routes = [
    // { path: "/products", component: productPage },
    // { path: "/products/print", component: printPage },
]

const router = new VueRouter({
    routes
})

var app;
window.addEventListener('DOMContentLoaded', (event) => {
    app = new Vue({
        router: router,
        el: '#app',
        data: {},
        computed: {}
    })
})