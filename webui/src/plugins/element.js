import Vue from 'vue'
import { Button,Menu,MenuItem,Radio,Message } from 'element-ui'

Vue.use(Button)
Vue.use(Menu)
Vue.use(MenuItem)
Vue.use(Radio)
Vue.component(Message.name, Message)
Vue.prototype.$message = Message;
