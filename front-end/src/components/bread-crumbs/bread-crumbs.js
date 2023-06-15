import {Link} from "react-router-dom";
import styles from "./bread-crumbs.module.css";

function BreadCrumbs({breadcrubms}) {
    return (
        <div>
            <h1><Link to={"/dashboard/admin"}>Dashboard</Link> / {breadcrubms.title}</h1>

            <div className={styles.fastActions}>
                {
                    breadcrubms.fastActions.map(action => (
                        <Link to={action.link}>{action.title}</Link>
                    ))
                }
            </div>
        </div>
    )
}

export default BreadCrumbs