import styles from "./admin.module.css"
import {logout} from "../../store/application"
import {useDispatch, useSelector} from "react-redux";
import {Link, Navigate} from "react-router-dom";
import {useTranslation} from "react-i18next";

function Admin(){
    let dispatch = useDispatch()
    let {t} = useTranslation()
    const application = useSelector((state) => state.application);

    return (
        <div className={styles.dashboard}>
            {!application.user.authorized &&
                <Navigate to="/auth/login"/>
            }

            <h1>{t("dashboard.header")}</h1>
            <p>{t("dashboard.hi")}, <b>{application.user.email}</b> {t("dashboard.intro")} 🧨</p>
            <div className={styles.fastActions}>
                <a onClick={(e) => {dispatch(logout())}}>{t("dashboard.fastActions.logout")}</a>
                <Link to={"/dashboard/admin/posts/create"}>{t("dashboard.fastActions.create_post")}</Link>
            </div>

            <div className={styles.dashboardBlocks}>
                <Link to="/dashboard/admin/posts" className={styles.dashboardBlock}>
                    <h1>{t("dashboard.blocks.posts.name")} 📋</h1>
                    <p>{t("dashboard.blocks.posts.desc")}</p>
                </Link>
                <Link to="/dashboard/admin/categories" className={styles.dashboardBlock}>
                    <h1>{t("dashboard.blocks.categories.name")} 🔗</h1>
                    <p>{t("dashboard.blocks.posts.desc")}</p>
                </Link>
                {/*<Link to="/dashboard/admin/tags" className={styles.dashboardBlock}>*/}
                {/*    <h1>Tags 🔋</h1>*/}
                {/*    <p>Create / modify / remove</p>*/}
                {/*</Link>*/}
            </div>
        </div>
    )
}

export default Admin