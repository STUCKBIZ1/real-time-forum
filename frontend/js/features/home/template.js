export function homeTemplate(){

    return `

    <div class="home">


        <header class="navbar">

            <div class="logo">
                Real Forum
            </div>


            <nav>

                <button id="profile-btn">
                    Profile
                </button>


                <button id="logout-btn">
                    Logout
                </button>

            </nav>


        </header>



        <div class="home-layout">



            <!-- LEFT SIDEBAR -->

            <aside id="left-sidebar" class="sidebar card">


                <div class="sidebar-item">
                    Home
                </div>


                <div class="sidebar-item">
                    My Posts
                </div>


                <div class="sidebar-item">
                    Categories
                </div>


                <div class="sidebar-item">
                    Saved
                </div>


            </aside>





            <!-- MAIN FEED -->

            <main id="feed">



                <!-- CREATE POST -->

                <section id="create-post" class="card">


                    <h3>
                        Create Post
                    </h3>


                    <textarea 
                        placeholder="What are you thinking?">
                    </textarea>


                    <button class="btn btn-primary">
                        Publish
                    </button>


                </section>





                <!-- POSTS -->

                <section id="posts-container">


                    <!-- posts will be injected here -->


                </section>




                <!-- LOADER -->

                <div id="posts-loader">


                </div>



            </main>






            <!-- RIGHT SIDEBAR -->


            <aside id="right-sidebar" class="sidebar card">



                <h3>
                    Online Users
                </h3>


                <div id="online-users">


                    <!-- users from websocket -->


                </div>



            </aside>



        </div>


    </div>


    `;

}