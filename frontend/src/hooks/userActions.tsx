import jwtDecode from 'jwt-decode';
import { useState, useEffect } from 'react';
import { IUserCookie } from './index.d';

export default function useUserActions() {
    const [user, setUser] = useState<IUserCookie>({"username" : "", exp: ""});
    const [isLoading, setLoading] = useState(true);
    const [token, setToken] = useState("");

    useEffect(() => {
        if (token !== "") {
            let decoded = jwtDecode<IUserCookie>(token);
            setUser(decoded);
        }
    }, [token]);
    
    return {
        user,
        setUser,
        isLoading,
        setLoading,
        token,
        setToken
    }
}