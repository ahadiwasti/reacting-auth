import axios from 'axios'
import { Toast } from 'primereact/toast';
import store from '@/store'

// create an axios instance
const service = axios.create({
    baseURL: process.env['REACTING_URL'],
    timeout: 15000, // request timeout,
})

// request interceptor
service.interceptors.request.use(
    config => {
        // Do something before request is sent
        const token = store.getters.token
        if (token) {
            config.headers['Authorization'] = 'Bearer ' + token
            // config.headers['Accept'] ='application/json'
            // config.headers['Content-Type'] = 'application/json'
        }
        return config
    },
    error => {
        // console.log(error)
        Promise.reject(error)
    }
)

// response interceptor
service.interceptors.response.use(
    response => {
        const res = response.data
        console.log(res)
        return
        if (res.code !== 200) {
            // Message({
            //     message: res.msg || 'Please contact the web admin',
            //     type: 'error',
            //     duration: 2 * 1000
            // })

            // // 50008:illegal token; 50012:Other clients logged in;  50014:Token expired;
            // if (
            //     res.code === 50008 ||
            //     res.code === 50012 ||
            //     res.code === 50014
            // ) {
            //     store.dispatch('FedLogOut').then(() => {
            //         location.reload()
            //     })
            // }
            // if (res.code === 10012 || res.code === 10015 || res.code === 401) {
            //     store.dispatch('FedLogOut').then(() => {
            //         location.reload()
            //     })
            // }
            return Promise.reject(response.data)
        } else {
            return response.data
        }
    },
    error => {
        const msg = error.response.data.message || error.message
        console.log(msg)
        return Promise.reject(error)
    }
)

export default service
