import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify';
import VueRouter from 'vue-router';
import axios from "axios";

Vue.config.productionTip = false
Vue.use(VueRouter)

function addHTTP(url) {
  if (!/^(?:f|ht)tps?:\/\//.test(url)) {
      url = "http://" + url;
  }
  return url;
}

const routes = [
  { path: '/:id', redirect: to => {
    const { _hash, params, _query } = to // eslint-disable-line no-unused-vars
    axios.get("http://localhost:8080/api/v1/id/" + params.id).then((response) => {
      console.log(addHTTP(response.data))
      location.replace(addHTTP(response.data))
      return response.data
    });
  }}
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
