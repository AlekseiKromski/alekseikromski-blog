import "./app.css"
import {Routes, Route} from "react-router-dom"
import Main from "./views/main/main"
import About from "./views/about/about"
import SideBar from "./components/sidebar/sidebar"

function App() {
    return (
        <div className="App">
            <SideBar/>
            <div className="content fontRoboto">
                <Routes>
                    <Route path="/" element={ <Main/> } />
                    <Route path="/about" element={ <About/> } />
                </Routes>
            </div>
        </div>
    )
}
export default App;
