import 'core-js/stable';
import 'regenerator-runtime/runtime';
import Vue from 'vue';
import App from './App.vue';
import store from './store'
import * as Wails from '@wailsapp/runtime';
import VTooltip from 'v-tooltip'

Vue.config.productionTip = false;
Vue.config.devtools = true;

Vue.use(VTooltip, { defaultDelay: 600, defaultOffset: 16 })

Wails.Init(() => {
	new Vue({
		render: h => h(App),
		store
	}).$mount('#app');
});
