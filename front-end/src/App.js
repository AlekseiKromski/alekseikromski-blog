import "./app.css"
import {Routes, Route} from "react-router-dom"
import Main from "./views/main/main"
import About from "./views/about/about"
import SideBar from "./components/sidebar/sidebar"
import Single from "./views/single/single";
import { Helmet } from 'react-helmet';
import {useDispatch} from "react-redux";
import {initAxios, setToken} from "./store/application"
import Auth from "./views/auth/auth";
import Admin from "./views/dashboard/admin";
import Posts from "./views/dashboard/posts/posts";
import PostCreate from "./views/dashboard/posts/create/create";

function App() {
    let dispatch = useDispatch()
    dispatch(initAxios())

    let account = JSON.parse(sessionStorage.getItem("account"))
    if (account != null && account != undefined) {
        dispatch(setToken({
            email: account.email,
            token: account.token
        }))
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
                    <Route path="/dashboard/admin/posts" element={ <Posts/> } />
                    <Route path="/dashboard/admin/posts/create" element={ <PostCreate/> } />
                </Routes>
            </div>
        </div>
    )
}
export default App;
