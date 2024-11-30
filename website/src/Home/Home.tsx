import {Box, Stack, Typography} from "@mui/material";
import Button from '@mui/material/Button';
import GitHubIcon from '@mui/icons-material/GitHub';
import {useNavigate} from "react-router-dom";
import DashboardIcon from '@mui/icons-material/Dashboard';

export const Home = () => {
    const navigate = useNavigate();

    return (
        <>
            <Box sx={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                height: '70vh',
                marginTop: '10vh',
                marginLeft: '10vw',
                marginRight: '10vw',
            }}>
                <Stack spacing={6} sx={{
                    display: 'flex',
                    justifyContent: 'center',
                    alignItems: 'center',
                }}>
                    <Box sx={{
                        width: '60vw',
                        display: 'flex',
                        justifyContent: 'center',
                        alignItems: 'center',
                        // borderColor: '#333945',
                        // borderStyle: 'solid',
                        // borderWidth: '2px',
                        borderRadius: '25px',
                        padding: '10px',
                        backgroundColor: '#343440',
                        boxShadow: '0px 0px 32px -2px rgba(0, 0, 0, 0.3)'
                    }}>
                        <Typography variant="h1" color="#f5f5f5" sx={{ letterSpacing: '2px' }}>
                            Markus Welscher
                        </Typography>
                    </Box>
                    <Box sx={{
                        display: 'flex',
                        justifyContent: 'center',
                        alignItems: 'center',
                        borderColor: '#333945',
                        borderStyle: 'solid',
                        borderWidth: '2px',
                        borderRadius: '25px',
                        padding: '10px',
                        width: '40vw',
                    }}>
                        <Typography variant="h3" color="#f5f5f5" sx={{ letterSpacing: '1px' }}>
                            Aspiring Software Engineer
                        </Typography>
                    </Box>

                    <Box sx={{
                        display: 'flex',
                        justifyContent: 'center',
                        alignItems: 'center',
                        padding: '20px',


                    }}>
                        <Stack spacing={8} direction="row">
                            <Button
                                variant="outlined"
                                startIcon={<GitHubIcon/>}
                                onClick={() => {
                                    window.location.href = 'https://github.com/WelscherMarkus';
                                }}
                                size="large"
                            >
                                Github
                            </Button>
                            <Button
                                variant="outlined"
                                startIcon={<img src={`${process.env.PUBLIC_URL}/LinkedIn.png`} alt="LinkedIn"
                                                style={{width: 23, height: 20}}/>}
                                onClick={() => {
                                    window.location.href = 'https://www.linkedin.com/in/markus-welscher-402599297/';
                                }}
                                size="large"
                            >
                                LinkedIn
                            </Button>
                            <Button
                                variant="outlined"
                                onClick={() => {
                                    navigate('/projects');
                                }}
                                startIcon={<DashboardIcon/>}
                                size="large"
                            >
                                Projects
                            </Button>
                        </Stack>
                    </Box>
                </Stack>
            </Box>
        </>
    );
};

export default Home;