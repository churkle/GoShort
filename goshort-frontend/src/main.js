import Vue from 'vue';
import App from './App.vue';
import vuetify from './plugins/vuetify';
import VueRouter from 'vue-router';
import axios from "axios";
import VueClipboard from 'vue-clipboard2';
import BlankPage from './components/BlankPage';

Vue.config.productionTip = false
Vue.use(VueRouter)
Vue.use(VueClipboard)

function addHTTP(url) {
  if (!/^(?:f|ht)tps?:\/\//.test(url)) {
      url = "http://" + url;
  }
  return url;
}

const routes = [
  { 
    path: '/:id', component: BlankPage, beforeEnter: async (to, from, next) => { // eslint-disable-line no-unused-vars
      const { _hash, params, _query } = to // eslint-disable-line no-unused-vars
      await axios.get("http://localhost:8080/api/v1/id/" + params.id).then((response) => {
        console.log(addHTTP(response.data))
        location.replace(addHTTP(response.data))
      });
    }
  },
  {
    path: '*',
    component: BlankPage
  }
]

const router = new VueRouter({
  routes,
  mode: 'history'
})

new Vue({
  router,
  vuetify,
  render: h => h(App)
}).$mount('#app')
