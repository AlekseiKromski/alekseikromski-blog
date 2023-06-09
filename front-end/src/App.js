import "./app.css"
import {Routes, Route} from "react-router-dom"
import Main from "./views/main/main"
import About from "./views/about/about"
import SideBar from "./components/sidebar/sidebar"
import Single from "./views/single/single";
import { Helmet } from 'react-helmet';

function App() {
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
                </Routes>
            </div>
        </div>
    )
}
export default App;
