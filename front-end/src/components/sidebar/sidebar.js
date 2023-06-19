import "./sidebar.css"
import CloseIcon from '@mui/icons-material/Close';
import MenuOpenIcon from '@mui/icons-material/MenuOpen';
import {useEffect, useState} from "react";
import {useDispatch, useSelector} from "react-redux";
import { importCategories } from '../../store/shared'
import {setDarkMode, setLang, setSideClosed} from '../../store/application'
import { useNavigate } from "react-router-dom";
import {FormControlLabel, Switch} from "@mui/material";
import {useTranslation} from "react-i18next";

function SideBar() {
    let navigate = useNavigate();
    let {i18n, t} = useTranslation()

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

    function switchLang(lang) {
        i18n.changeLanguage(lang)
        dispatch(setLang(lang))
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
                                <a onClick={() => red("/")}>{t("sidebar.posts")}</a>
                            </li>
                            <li>
                                <a onClick={(e) => red("/about")}>{t("sidebar.about")}</a>
                            </li>
                            {application.user.authorized &&
                                <li>
                                    <a onClick={(e) => red("/dashboard/admin")}>{t("sidebar.dashboard")} 🛠️</a>
                                </li>
                            }
                        </ul>

                        <div className="sidebarMenu">
                            <h1>{t("sidebar.categories")}</h1>

                            <a onClick={() => red("/")}>All</a>
                            {shared.categories.map(category => (
                                <a onClick={() => red(`/${category.id}`)} key={category.id}>{category.name}</a>
                            ))}
                        </div>

                        <div className="sidebarMenu">
                            <h1>{t("sidebar.settings")}</h1>
                            <div className="">
                                <FormControlLabel
                                    control={
                                        <Switch
                                            checked={application.darkMode}
                                            onChange={(event) => {
                                                dispatch(setDarkMode(event.target.checked))
                                            }}
                                            name="theme"
                                            color="default"
                                        />
                                    }
                                    label={t("sidebar.switch")}
                                />
                            </div>
                        </div>

                        <div className="sidebarMenu sidebarMenuLanguage">
                            <h1>{t("sidebar.language")}</h1>
                            <a
                                className={`${application.lang === "en" ? "active": "link"}`}
                                onClick={() => switchLang("en")}
                            >🇺🇸</a>
                            <a
                                className={`${application.lang === "et" ? "active": "link"}`}
                                onClick={() => switchLang("et")}
                            >🇪🇪</a>
                            <a
                                className={`${application.lang === "ru" ? "active": "link"}`}
                                onClick={() => switchLang("ru")}
                            >🇷🇺</a>
                        </div>
                    </div>
                    <div className="copyright">
                        <small>{t("sidebar.copyright")} Aleksei Kromski 2023</small>
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