import {Link} from "react-router-dom";
import styles from "./bread-crumbs.module.css";

function BreadCrumbs({breadcrubms}) {
    return (
        <div>
            <h1 className={styles.breadCrumbs}>

                {breadcrubms.links.map(link => (
                    <Link to={link.link}>{link.title} / </Link>
                ))}
                    <span>{breadcrubms.title}</span>

            </h1>

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