import "./sidebar.css"
import {Link} from "react-router-dom";
import CloseIcon from '@mui/icons-material/Close';
import MenuOpenIcon from '@mui/icons-material/MenuOpen';
import {useState} from "react";


function SideBar() {
    let [close, setClose] = useState(true);
    let closeFunction = () => {
        setClose(!close)
    }

    return (
        <div className="mainSideBar">
            <div className={`sideBarBlock ${close ? "sideBarMinimal-show" : "sideBar-hide"}`}>
                <div className="sideBar">
                    <div className="">
                        <h1 className="fontRighteous">
                            Blog
                            <CloseIcon
                                className="close"
                                onClick={closeFunction}
                            />
                        </h1>
                        <p>Small blog about my life</p>
                    </div>
                    <div className="links">
                        <ul>
                            <li>
                                <Link to="/">Posts</Link>
                            </li>
                            <li>
                                <Link to="/categories">Categories</Link>
                            </li>
                            <li>
                                <Link to="/links">Links</Link>
                            </li>
                            <li>
                                <Link to="/about">About</Link>
                            </li>
                        </ul>

                        <div className="copyright">
                            <small>Copyright Aleksei Kromski 2023</small>
                        </div>
                    </div>
                </div>
            </div>

            <div className={`sideBarMinimalBlock ${!close ? "sideBarMinimal-show" : "sideBar-hide"}`}>
                <MenuOpenIcon onClick={closeFunction} className="open"/>
            </div>
        </div>
    )
}

export default SideBar