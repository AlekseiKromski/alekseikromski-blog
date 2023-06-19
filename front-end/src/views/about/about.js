import { Link } from "react-router-dom";
import styles from "./about.module.css"
import FacebookIcon from '@mui/icons-material/Facebook';
import {useSelector} from "react-redux";
import {useTranslation} from "react-i18next";

function About() {
    const {t} = useTranslation()
    const application = useSelector((state) => state.application);

    return (
        <div className={`${styles.about} ${!application.sideClosed ? "static" : ""}`}>
            <div className={styles.aboutBlock}>
                <div className={styles.aboutImage}>
                    <img src={require("../../images/avatar.jpeg")} alt=""/>
                </div>
                <div className={styles.aboutText}>
                    <h1>{t("about.title")} 🚀</h1>
                    <p>
                        {t("about.description")}
                    </p>
                    <div className={styles.links}>
                        <Link to="/">{t("about.buttons.read_posts")} 👆</Link>
                        <a href="https://alekseikromski.com" target="_blank">{t("about.buttons.check_portfolio")} 🧨</a>
                    </div>
                    <div className={styles.links}>
                        <a href="https://www.facebook.com/aleksei.kromski.3/" target="_blank">
                            {t("about.buttons.fb")}
                            <FacebookIcon/>
                        </a>
                        <a href="mailto:aleskeikromski@outlook.com">{t("about.buttons.click_email")} ✉️</a>
                        <a href="https://alekseikromski.com/" target="_blank">{t("about.buttons.cv")} 📋️</a>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default About;