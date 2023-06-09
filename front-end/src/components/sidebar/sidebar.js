import "./sidebar.css"
import CloseIcon from '@mui/icons-material/Close';
import MenuOpenIcon from '@mui/icons-material/MenuOpen';
import {useEffect, useState} from "react";
import {useDispatch, useSelector} from "react-redux";
import { importCategories } from '../../store/shared'
import { setSideClosed } from '../../store/application'
import { useNavigate } from "react-router-dom";

function SideBar() {
    let navigate = useNavigate();

    //State
    let [isMobile, setIsMobile] = useState(false)

    window.addEventListener("resize", (e) => {
        checkSize()
    })

    //Redux
    const dispatch = useDispatch()
    const shared = useSelector((state) => state.shared);
    const application = useSelector((state) => state.application);

    let closeFunction = () => {
        console.log('SSS')
        dispatch(setSideClosed())
    }

    async function getCategories () {
        await application.axios.get("/v1/category/all").then(response => {
            dispatch(importCategories(response.data))
        })
    }

    function checkSize() {
        if (window.innerWidth <= 800){
            setIsMobile(true)
        }else {
            setIsMobile(false)
        }
    }

    function sideBarIdentify() {
        if (isMobile == false) {
            return application.sideClosed ? "sideBarMinimal-show" : "sideBar-hide"
        }
        return application.sideClosed ? "sideBar-hide" : "sideBarMinimal-show"
    }

    function sideBarMinimalIdentify() {
        if (isMobile == false) {
            return !application.sideClosed ? "sideBarMinimal-show" : "sideBar-hide"
        }
        return !application.sideClosed ? "sideBar-hide" : "sideBarMinimal-show"
    }

    function red(to) {
        if (isMobile) {
            dispatch(setSideClosed(true))
        }
        return navigate(to)
    }

    useEffect(() => {
        getCategories()
        checkSize()
    }, [])

    return (
        <div className="mainSideBar">
            <div className={`sideBarBlock ${sideBarIdentify()}`}>
                <div className="sideBar">
                    <div className="">
                        <h1 className="fontRighteous">
                            <a onClick={() => red("/")} className="logo">
                                <img src={require("../../images/logo.png")} alt=""/>
                                Blog
                            </a>
                            <CloseIcon
                                className="close"
                                onClick={() => {closeFunction()}}
                            />
                        </h1>
                    </div>
                    <div className="links">
                        <ul>
                            <li>
                                <a onClick={() => red("/")}>Posts</a>
                            </li>
                            <li>
                                <a onClick={(e) => red("/about")}>About</a>
                            </li>
                            {application.user.authorized &&
                                <li>
                                    <a onClick={(e) => red("/dashboard/admin")}>Dashboard 🛠️</a>
                                </li>
                            }
                        </ul>

                        <div className="categories">
                            <h1>Categories</h1>

                            <a onClick={() => red("/")}>All</a>
                            {shared.categories.map(category => (
                                <a onClick={() => red(`/${category.ID}`)} key={category.id}>{category.name}</a>
                            ))}
                        </div>
                    </div>
                    <div className="copyright">
                        <small>Copyright Aleksei Kromski 2023</small>
                    </div>
                </div>
            </div>

            <div className={`sideBarMinimalBlock ${sideBarMinimalIdentify()}`}>
                <MenuOpenIcon onClick={closeFunction} className="open"/>
            </div>
        </div>
    )
}

export default SideBar