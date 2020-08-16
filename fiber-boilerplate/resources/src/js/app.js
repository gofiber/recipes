import Vue from 'vue'
import App from './App.vue'
import Router from 'vue-router'
import {VuejsDatatableFactory} from 'vuejs-datatable';

import {FontAwesomeIcon} from '@fortawesome/vue-fontawesome'

import store from './store'
import routes from './routes'
import "./icons"

Vue.component('FontAwesome', FontAwesomeIcon)
Vue.config.productionTip = false

Vue.use(Router)

VuejsDatatableFactory.useDefaultType(false)
    .registerTableType('datatable', tableType => tableType.mergeSettings({
        table: {class: 'datatable table-fixed w-full'},
    }));
Vue.use(VuejsDatatableFactory);

const router = new Router({
    mode: 'hash',
    routes
})
new Vue({
    render: h => h(App),
    router,
    store
}).$mount('#app')

