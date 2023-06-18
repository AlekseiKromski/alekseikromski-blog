import styles from "./admin.module.css"
import {logout} from "../../store/application"
import {useDispatch, useSelector} from "react-redux";
import {Link, Navigate} from "react-router-dom";

function Admin(){
    let dispatch = useDispatch()
    const application = useSelector((state) => state.application);

    return (
        <div className={styles.dashboard}>
            {!application.user.authorized &&
                <Navigate to="/auth/login"/>
            }

            <h1>Dashboard</h1>
            <p>Hi, <b>{application.user.email}</b> this is your dashboard 🧨</p>
            <div className={styles.fastActions}>
                <a onClick={(e) => {dispatch(logout())}}>logout</a>
                <Link to={"/dashboard/admin/posts/genericCU"}>Create post</Link>
            </div>

            <div className={styles.dashboardBlocks}>
                <Link to="/dashboard/admin/posts" className={styles.dashboardBlock}>
                    <h1>Posts 📋</h1>
                    <p>Create / modify / remove / manage comments</p>
                </Link>
                <Link to="/dashboard/admin/categories" className={styles.dashboardBlock}>
                    <h1>Categories 🔗</h1>
                    <p>Create / modify / remove</p>
                </Link>
                <Link to="/dashboard/admin/tags" className={styles.dashboardBlock}>
                    <h1>Tags 🔋</h1>
                    <p>Create / modify / remove</p>
                </Link>
            </div>
        </div>
    )
}

export default Admin