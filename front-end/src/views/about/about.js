import { Link } from "react-router-dom";

function About() {
    return (
        <div>
            <h1>Something about</h1>
            <Link to="/">Click to view our main page</Link>
        </div>
    );
}

export default About;