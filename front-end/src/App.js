import "./app.css"
import {Routes, Route, Link, useNavigate} from "react-router-dom"
import Main from "./views/main/main"
import About from "./views/about/about"
import SideBar from "./components/sidebar/sidebar"
import Single from "./views/single/single";
import { Helmet } from 'react-helmet';
import {useDispatch, useSelector} from "react-redux";
import {initAxios, setToken} from "./store/application"
import Auth from "./views/auth/auth";
import Admin from "./views/dashboard/admin";
import Generic from "./views/dashboard/generic/generic";
import PostCreate from "./views/dashboard/generic/post/create/create";
import PostEdit from "./views/dashboard/generic/post/edit/edit";
import {useEffect, useState} from "react";
import styles from "./views/dashboard/generic/generic.module.css";

function App() {

    const navigate = useNavigate()
    let dispatch = useDispatch()
    dispatch(initAxios())


    useEffect(() => {
        let account = JSON.parse(sessionStorage.getItem("account"))
        if (account != null) {
            dispatch(setToken({
                email: account.email,
                token: account.token
            }))
        }
    }, [])


    let generic = {
        categories: {
            breadcrumbs: {
                title: "Categories",
                fastActions: [
                    {
                        title: "Create category",
                        link: "/dashboard/admin/categories/create"
                    }
                ]
            },
            data: {
                ids: ["ID", "Name", "Action"],
                handler: (cb, application) => {
                    application.axios.get("/v1/category/all").catch(
                        cb([])
                    ).then(response => {
                        if (response.data != null) {
                            cb(response.data)
                            return
                        }
                        cb([])
                    })
                },
                renderTable(cat, style, actions){
                    //this -> this data object
                    return (
                        <tr>
                            <th>{cat.id}</th>
                            <th className={style.title}>
                                {cat.name}
                            </th>
                            <th>
                                <div className={styles.action}>
                                    {
                                        actions.map(action => (
                                            <button onClick={() => action.func()}>{action.title}</button>
                                        ))
                                    }
                                </div>
                            </th>
                        </tr>
                    )
                },
                urls: {
                    edit: "/dashboard/admin/categories/edit",
                    delete: "/v1/category/delete"
                }
            }
        },
        posts: {
            breadcrumbs: {
                title: "Posts",
                fastActions: [
                    {
                        title: "Create post",
                        link: "/dashboard/admin/posts/create"
                    }
                ]
            },
            data: {
                ids: ["ID", "Name", "Action"],
                handler: (cb, application) => {
                    application.axios.get("/v1/post/get-last-posts/15/0").catch(
                        cb([])
                    ).then(response => {
                        if (response.data != null) {
                            cb(response.data)
                            return
                        }
                        cb([])
                    })
                },
                renderTable(post, style, actions){
                    //this -> this data object
                    return (
                        <tr>
                            <th>{post.id}</th>
                            <th className={style.title}>
                                <Link to={"/post/" + post.id}>{post.title}</Link>
                            </th>
                            <th>
                                <div className={styles.action}>
                                    {
                                        actions.map(action => (
                                            <button onClick={() => action.func()}>{action.title}</button>
                                        ))
                                    }
                                </div>
                            </th>
                        </tr>
                    )
                },
                urls: {
                    edit: "/dashboard/admin/posts/edit",
                    delete: "/v1/post/delete"
                }
            }
        }
    }

    return (
        <div className="App">
            <Helmet>
                <title>Blog | Aleksei Kromski</title>
            </Helmet>
            <SideBar/>
            <div className="content fontRoboto">
                <Routes>
                    <Route path="/:categoryID?" element={ <Main/> } />
                    <Route path="/post/:id" element={ <Single/> } />
                    <Route path="/about" element={ <About/> } />
                    <Route path="/auth/login" element={ <Auth/> } />

                    <Route path="/dashboard/admin" element={ <Admin/> } />
                    <Route path="/dashboard/admin/posts" element={ <Generic
                        settings={generic.posts}
                    /> } />
                    <Route path="/dashboard/admin/posts/create" element={ <PostCreate/> } />
                    <Route path="/dashboard/admin/posts/edit/:id" element={ <PostEdit/> } />

                    <Route path="/dashboard/admin/categories" element={ <Generic
                        settings={generic.categories}
                    /> } />

                </Routes>
            </div>
        </div>
    )
}
export default App;
