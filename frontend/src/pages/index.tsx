import type {GetServerSideProps, NextPage} from "next";
import Image from "next/image";
import {urqlClient} from "@ka-libs/gql-requests";
import styles from "@ka-styles/Home.module.css";
import {GetUsersDocument, User} from "@ka-graphql/generated.graphql";
import {Box, Stack} from "@mui/system";
import {Avatar, List, ListItem, ListItemAvatar, ListItemText} from "@mui/material";

type Props = {
    users: User[]
};

const Home: NextPage<Props> = (props) => {
    return (
        <Stack
            sx={{
                minHeight: "100vh",
            }}
        >
            <List sx={{width: "100%", maxWidth: 360, bgcolor: "background.paper"}}>
                {props.users.map((user) => (
                    <ListItem key={user.id}>
                        <ListItemAvatar>
                            <Avatar>絵</Avatar>
                        </ListItemAvatar>
                        <ListItemText primary={user.nickname} secondary="公開日"/>
                    </ListItem>
                ))}
            </List>
            <Box
                sx={{
                    bgColor: "palette.primary.dark",
                    backgroundColor: (theme) => theme.palette.primary.dark,
                    color: (theme) =>
                        theme.palette.getContrastText(theme.palette.primary.dark),
                    py: 3,
                    textAlign: "center",
                    marginTop: "auto",
                }}
            >

                <footer>
                    <a
                        href="https://vercel.com?utm_source=create-next-app&utm_medium=default-template&utm_campaign=create-next-app"
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        Powered by{" "}
                        <span className={styles.logo}>
                <Image src="/vercel.svg" alt="Vercel Logo" width={72} height={16}/>
              </span>
                    </a>
                </footer>
            </Box>
        </Stack>
    );
};

export const getServerSideProps: GetServerSideProps<Props> = async () => {
    try {
        const client = await urqlClient();

        const result = await client.query(GetUsersDocument, {}).toPromise();
        console.log(result)
        if (!result.data || !result.data.users) {
            console.error('No data returned from GraphQL query');
            return {notFound: true};
        }

        return {
            props: {
                users: result.data.users,
            },
        };
    } catch (e) {
        // @ts-ignore
        console.error('Error during fetch data:', e.message);
        return {
            props: {
                // @ts-ignore
                error: e.message
            },
            notFound: true,
        };
    }
};

export default Home;