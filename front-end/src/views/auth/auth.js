import styles from "./auth.module.css"
import {Link, Navigate, useNavigate} from "react-router-dom"
import {useDispatch, useSelector} from "react-redux";
import {useState} from "react";
import {setToken} from "../../store/application"
import {useTranslation} from "react-i18next";

function Auth(){
    const {t} = useTranslation()
    let navigate = useNavigate()
    let [email, setEmail] = useState("")
    let [password, setPassword] = useState("")
    let [error, setError] = useState("")

    //redux
    const dispatch = useDispatch()
    const application = useSelector((state) => state.application);

    function login() {
        if (email != "" && password != "") {
            application.axios.post("/v1/login", {
                email: email,
                password: password
            }).then(response => {
                if (response.data != "") {
                    dispatch(setToken({
                        email: email,
                        token: response.data
                    }))
                    navigate("/dashboard/admin");
                    return
                }

                setError("cannot get token")
            }).catch(e => {
                if (e.response !== undefined) {
                    setError(e.response.data.message)
                }
                console.error(e)
            })
        }
    }

    return (
        <div className={`${styles.login} ${!application.sideClosed ? "static" : ""}`}>
            <div className={styles.loginBlock}>
                <h1>{t("auth.login")} 🚀</h1>

                {error != "" &&
                    <p>{error}</p>
                }

                {application.user.authorized &&
                    <Navigate to="/dashboard/admin"/>
                }

                <label htmlFor="email">{t("auth.email")}</label>
                <input name="email" type="text" onChange={(e) => setEmail(e.target.value)} value={email}/>

                <label htmlFor="password">{t("auth.password")}</label>
                <input name="password" type="text" onChange={(e) => setPassword(e.target.value)} value={password}/>

                <button onClick={(e) => login()}>{t("auth.signin")}</button>
                <Link to="/">{t("auth.back")}</Link>
            </div>
        </div>
    )
}

export default Auth