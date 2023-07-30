import axios from 'axios'

axios.defaults.headers.post['Content-Type'] = 'application/x-www-form-urlencoded';
/* axios.defaults.headers.post['Content-Type'] = 'application/json;charset=UTF-8'; */
axios.defaults.headers.get['Content-Type'] = 'application/x-www-form-urlencoded';
axios.defaults.transformRequest = [function (data) {
    let src = ''
    for (let item in data) {
      src += encodeURIComponent(item) + '=' + encodeURIComponent(data[item]) + '&'
    }
    return src
}]


export function getResult(param){
    return axios.request({
        method: "GET",
        url: "/api/user/search",
        params : param
    })
}

export function postCreate(param){
    return axios.request({
        method: "POST",
        url: "/api/user/create",
        params : param
    })
}

export function getCompanies(param){
    return axios.request({
        method: "GET",
        url: "/api/util/company",
        params : param
    })
}

export function postSignup(param){
    return axios.request({
        method: "POST",
        url: "/api/user/signup",
        params : param
    })
}

export function postLogin(param){
    return axios.request({
        method: "POST",
        url: "/api/user/login",
        data : param
    })
}

export function postAtc(param){
    return axios.request({
        method: "POST",
        url: "/api/publisher/postatc",
        data : param
    })
}

export function postMultiAtc(param){
    return axios.request({
        method: "POST",
        url: "/api/publisher/postmultiatc",
        data : param
    })
}

export function getAtc(param){
    return axios.request({
        method: "GET",
        url: "/api/publisher/getatc",
        params: param
    })
}

export function getCompanyOptions(param){
    console.log(param)
    return axios.request({
        method: "GET",
        //url: "/api/util/getcompanies",
        url: "/api/util/company",
        params:param
    })
}

export function getPublisherOptions(param){
    return axios.request({
        method: "GET",
        url: "/api/util/getcheckpublishers",
        params:param
    })
}


export function getCertificate(param){
    return axios.request({
        method: "GET",
        url: "/api/util/getcert",
        params:param
    })
}


export function getPrivateKey(param){
    return axios.request({
        method: "GET",
        url: "/api/util/getpriv",
        params:param
    })
}

export function postEdit(param){
    return axios.request({
        method: "POST",
        url: "/api/publisher/postedit",
        params:param
    })
}

export function postDelete(param){
    return axios.request({
        method: "POST",
        url: "/api/publisher/postdelete",
        params:param
    })
}

export function getArns(param){
    return axios.request({
        method: "GET",
        url: "/api/publisher/getarns",
        params: param
    })
}

export function getRoutes(param){
    return axios.request({
        method: "GET",
        url: "/api/publisher/getroutes",
        params: param
    })
}

export function getAuths(param){
    return axios.request({
        method: "GET",
        url: "/api/publisher/getauths",
        params: param
    })
}

export function getLinks(param){
    return axios.request({
        method: "GET",
        url: "/api/publisher/getlinks",
        params: param
    })
}

export function getLinkOptions(param){
    return axios.request({
        method: "GET",
        url: "/api/publisher/getlinkoptions",
        params: param
    })
}

export function postRegLink(param){
    return axios.request({
        method: "POST",
        url: "/api/publisher/postreglink",
        data : param
    })
}

export function getRegistLink(param){
    return axios.request({
        method: "GET",
        url: "/api/publisher/getreglink",
        params : param
    })
}