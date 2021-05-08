import { useContext } from "react";
import S3UploadInput from "../components/S3UploadInput";
import { UserContext } from "../hooks/UserContext";

const Dashboard = () => {
    const {setToken} = useContext(UserContext);
    return (
        <div>
            <span>
                <h1>Dashboard</h1>
                <button onClick={() => {setToken("invalid")}}>Logout</button>
            </span>
            <S3UploadInput/>
        </div>
    );
}

export default Dashboard;