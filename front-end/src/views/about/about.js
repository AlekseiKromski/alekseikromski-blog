import { Link } from "react-router-dom";
import styles from "./about.module.css"
import FacebookIcon from '@mui/icons-material/Facebook';
import {useSelector} from "react-redux";

function About() {
    const application = useSelector((state) => state.application);

    return (
        <div className={`${styles.about} ${!application.sideClosed ? "static" : ""}`}>
            <div className={styles.aboutBlock}>
                <div className={styles.aboutImage}>
                    <img src={require("../../images/avatar.jpeg")} alt=""/>
                </div>
                <div className={styles.aboutText}>
                    <h1>Hi, my name is Aleksei Kromski 🚀</h1>
                    <p>
                        Here you can find small posts about my IT life and the progress of various projects, as well as my thoughts on some technology.
                    </p>
                    <div className={styles.links}>
                        <Link to="/">Read posts 👆</Link>
                        <a href="https://alekseikromski.com" target="_blank">Check portfolio 🧨</a>
                    </div>
                    <div className={styles.links}>
                        <a href="https://www.facebook.com/aleksei.kromski.3/" target="_blank">
                            Facebook
                            <FacebookIcon/>
                        </a>
                        <a href="mailto:aleskeikromski@outlook.com">Click to mail ✉️</a>
                        <a href="https://alekseikromski.com/" target="_blank">Get CV 📋️</a>
                    </div>
                </div>
            </div>
        </div>
    );
}

export default About;