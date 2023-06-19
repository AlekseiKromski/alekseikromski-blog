import "./app.css"
import {Routes, Route, Link, useNavigate} from "react-router-dom"
import Main from "./views/main/main"
import About from "./views/about/about"
import SideBar from "./components/sidebar/sidebar"
import Single from "./views/single/single";
import { Helmet } from 'react-helmet';
import {useDispatch, useSelector} from "react-redux";
import {setDarkMode, setToken} from "./store/application"
import Auth from "./views/auth/auth";
import Admin from "./views/dashboard/admin";
import Generic from "./views/dashboard/generic/generic";
import PostCreate from "./views/dashboard/generic/post/create/create";
import PostEdit from "./views/dashboard/generic/post/edit/edit";
import {useEffect} from "react";
import styles from "./views/dashboard/generic/generic.module.css";
import GenericCU from "./views/dashboard/generic/genericCU/genericCU";
import {importCategories} from "./store/shared";
import {useTranslation} from "react-i18next";

function App() {
    let navigate = useNavigate()
    let dispatch = useDispatch()
    let {t} = useTranslation()

    const application = useSelector((state) => state.application);

    useEffect(() => {
        let account = JSON.parse(sessionStorage.getItem("account"))
        if (account != null) {
            dispatch(setToken({
                email: account.email,
                token: account.token
            }))
        }

        let darkMode = sessionStorage.getItem("darkMode")
        if (darkMode != null && darkMode === "true") {
            dispatch(setDarkMode(true))
        } else {
            dispatch(setDarkMode(false))
        }
    }, [])


    let generic = {
        categories: {
            breadcrumbs: {
                title: t("generic.categories.breadcumbs_title"),
                links: [
                    {
                        title: t("generic.categories.breadcumbs_link_title"),
                        link: "/dashboard/admin"
                    },
                ],
                fastActions: [
                    {
                        title: t("generic.categories.breadcumbs_fastActions_title"),
                        link: "/dashboard/admin/categories/create"
                    }
                ]
            },
            data: {
                ids: [t("generic.categories.table_id"), t("generic.categories.table_name"), t("generic.categories.table_action")],
                handler: (cb, application) => {
                    application.axios.get("/v1/category/all").catch(
                        cb([])
                    ).then(response => {
                        if (response.data != null) {
                            cb(response.data)
                            dispatch(importCategories(response.data))
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
            },
            create: {
                buttonName: t("generic.categories.create_button_name"),
                breadcrumbs: {
                    title: t("generic.categories.create_breadcumbs_title"),
                    links: [
                        {
                            title: t("generic.categories.breadcumbs_link_title"),
                            link: "/dashboard/admin"
                        },
                        {
                            title: t("generic.categories.breadcumbs_title"),
                            link: "/dashboard/admin/categories"
                        },
                    ],
                    fastActions: []
                },
                func: (data, application) => {
                    application.axios.post("v1/category/create", {
                        name: data
                    }).then( () => {
                        navigate("/dashboard/admin/categories")
                    })
                },
                preFunc: () => {}
            },
            update: {
                buttonName: t("generic.categories.update_button_name"),
                breadcrumbs: {
                    title: t("generic.categories.update_breadcumbs_title"),
                    links: [
                        {
                            title: t("generic.categories.breadcumbs_link_title"),
                            link: "/dashboard/admin"
                        },
                        {
                            title: t("generic.categories.breadcumbs_title"),
                            link: "/dashboard/admin/categories"
                        },
                    ],
                    fastActions: []
                },
                func: (id, data, application) => {
                    application.axios.post("v1/category/edit/", {
                        id: Number.parseInt(id),
                        name: data
                    }).then( () => {
                        navigate("/dashboard/admin/categories")
                    })
                },
                preFunc: (id, shared, application) => {
                    let cat = shared.categories.find( cat => {
                        if (cat.id == id) {
                            return cat
                        }
                    })
                    if (cat == undefined) {
                        return ""
                    }
                    return cat.name
                }
            }
        },
        posts: {
            breadcrumbs: {
                title: t("generic.posts.breadcumbs_title"),
                links: [
                    {
                        title: t("generic.posts.breadcumbs_link_title"),
                        link: "/dashboard/admin"
                    },
                ],
                fastActions: [
                    {
                        title: t("generic.posts.breadcumbs_fastActions_title"),
                        link: "/dashboard/admin/posts/create"
                    }
                ]
            },
            data: {
                ids: [t("generic.posts.table_id"), t("generic.posts.table_name"), t("generic.posts.table_action")],
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
                                <div className={style.action}>
                                    {
                                        actions.map(action => (
                                            <button key={action.title} onClick={() => action.func()}>{action.title}</button>
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
        <div className={`App ${application.darkMode ? "dark" : ""}`}>
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
                    <Route path="/dashboard/admin/categories/create" element={ <GenericCU
                        settings={generic.categories.create}
                    /> } />
                    <Route path="/dashboard/admin/categories/edit/:id" element={ <GenericCU
                        settings={generic.categories.update}
                    /> } />
                </Routes>
            </div>
        </div>
    )
}
export default App;
