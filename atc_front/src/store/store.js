import Vue from 'vue'
import Vuex from 'vuex'
import Cookie from 'vue-cookies'

Vue.use(Vuex)

export default new Vuex.Store({
    state:{
        userid: Cookie.get("userid"),
        usercompany: Cookie.get("usercompany"),
        token: Cookie.get("token"),
    },
    mutations:{
        saveToken: function(state , userToken){
            state.userid = userToken.userid;
            state.usercompany = userToken.usercompany;
            state.cncompany = userToken.cncompany;
            state.token = userToken.token;
            Cookie.set("userid", userToken.userid, "7d");
            Cookie.set("usercompany", userToken.usercompany, "7d");
            Cookie.set("cncompany", userToken.cncompany, "7d");
            Cookie.set("token", userToken.token, "7d");
        },
        clearToken: function(state){
            state.userid = null;
            state.usercompany = null;
            state.cncompany = null;
            state.token = null;
            Cookie.remove('userid');
            Cookie.remove('usercompany');
            Cookie.remove('cncompany');
            Cookie.remove('token');
        }
    }
})