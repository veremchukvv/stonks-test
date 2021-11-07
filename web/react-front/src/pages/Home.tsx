import React, { useEffect, useState } from 'react';

const Home = () => {
    const [name, setName] = useState('')

    useEffect(() => {
        (
            async () => {
                const response = await fetch('http://localhost:8000/users/user', {
                    headers: {'Content-Type':'application/json'},
                    credentials: 'include',
                    })
                const content = await response.json()

                setName(content.name)

                }
        )()
    })

    return (
        <div>
            {name ? 'Welcome ' + name : 'Please sign in' }
        </div>
    );
};

export default Home;