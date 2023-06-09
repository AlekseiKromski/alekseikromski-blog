import "./sidebar.css"
import {Link} from "react-router-dom";
import CloseIcon from '@mui/icons-material/Close';
import MenuOpenIcon from '@mui/icons-material/MenuOpen';
import {useEffect, useState} from "react";
import axios from "axios";
import {useDispatch, useSelector} from "react-redux";
import { importCategories } from '../../store/shared'

function SideBar() {
    //Redux
    const dispatch = useDispatch()
    const shared = useSelector((state) => state.shared);

    let [close, setClose] = useState(true);
    let closeFunction = () => {
        setClose(!close)
    }

    async function getCategories () {
        await axios.get("http://localhost:3001/v1/category/all").then(response => {
            dispatch(importCategories(response.data))
        })
    }

    useEffect(() => {
        getCategories()
    }, [])

    return (
        <div className="mainSideBar">
            <div className={`sideBarBlock ${close ? "sideBarMinimal-show" : "sideBar-hide"}`}>
                <div className="sideBar">
                    <div className="">
                        <h1 className="fontRighteous">
                            <Link to="/" className="logo">
                                <img src={require("../../images/logo.png")} alt=""/>
                                Blog
                            </Link>
                            <CloseIcon
                                className="close"
                                onClick={closeFunction}
                            />
                        </h1>
                    </div>
                    <div className="links">
                        <ul>
                            <li>
                                <Link to="/">Posts</Link>
                            </li>
                            <li>
                                <Link to="/about">About</Link>
                            </li>
                        </ul>

                        <div className="categories">
                            <h1>Categories</h1>

                            <Link to={`/`}>All</Link>
                            {shared.categories.map(category => (
                                <Link to={`/${category.ID}`} key={category.id}>{category.name}</Link>
                            ))}
                        </div>
                    </div>
                    <div className="copyright">
                        <small>Copyright Aleksei Kromski 2023</small>
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