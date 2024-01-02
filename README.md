### Check Commit Logs
<img width="447" alt="Screenshot 2023-12-29 141249" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/203c0ca6-868d-4277-9700-21d9b2e3a674">
<img width="629" alt="Screenshot 2023-12-29 141414" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/2108d1ed-71a4-4e72-b9b5-57cd88b8128c">

### Change 'pick' to 'reword' to change the commit message 
<img width="570" alt="Screenshot 2023-12-29 141541" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/97f4a5a8-6e83-4671-bf20-5305167b1e1a">
<img width="504" alt="Screenshot 2023-12-29 141615" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/876ff70d-ba73-4965-b118-3d8d2d3a7555">

### Change Commit message here 
<img width="471" alt="Screenshot 2023-12-29 141700" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/40bdd3dd-2abc-4180-8e2a-2a5273ace5ca">
<img width="484" alt="Screenshot 2023-12-29 141728" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/2b55cf10-6cb3-4438-8585-272072721043">
<img width="486" alt="Screenshot 2023-12-29 141754" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/1976281c-4dab-48d8-aeb7-02a247597a51">

### Change 'pick' to 'squash' to combine the commits

Lets say you have commit logs like this (for example)
1. pick <commit_id_1> <commit_message_1>
2. pick <commit_id_2> <commit_message_2>
3. pick <commit_id_3> <commit_message_3>

You want to combine commit 2 and commit 3 here , then write 'squash' inplace of 'pick' in the commit 3 to combine combine 2 and 3 into one .  

<img width="604" alt="Screenshot 2023-12-29 142015" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/b68c3b27-45b0-4d8b-8b49-6dd04de7f5de">
<img width="476" alt="Screenshot 2023-12-29 142043" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/9bbedf04-baae-4021-9d1d-1c8296880c48">

### Modify the combined commit message here

<img width="458" alt="Screenshot 2023-12-29 142148" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/c2f61107-208b-4c59-8111-8252d8ade2f5">
<img width="506" alt="Screenshot 2023-12-29 142226" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/41fd2110-7471-4bea-bc2a-10146badb839">

Now lets say you want to make the local changes and main github account changes in sync you cannot push changes normally . If you push the changes normally then
this kind of problem will be seen . 


<img width="584" alt="Screenshot 2023-12-29 144130" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/ef34e93c-cd44-4006-95b8-ceae6823172c">

To make sync with main github acount you have to force push the changes .

<img width="578" alt="Screenshot 2023-12-29 143655" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/bda0ab54-ae73-4dce-bb37-c52860b628ec">

### Cherrypick

Now lets say you have made some changes on the local repo , but you have commited the changes to the wrong branch of repository . Now to do solve this issue , cherrypick 
concept is used .

<img width="518" alt="Screenshot 2023-12-29 150347" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/77303e82-a54f-4408-886d-d03dab5a0233">

Copy the commit id .

<img width="481" alt="Screenshot 2023-12-29 150533" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/4ff94d4b-9448-4014-9f17-07a1be88103c">

Then switch to the desired branch(here i have changed to main branch) and write the command 
```
git cherrypick <commit-id>

```

<img width="468" alt="Screenshot 2023-12-29 150753" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/920800b7-8611-4c86-b053-f6dd818a2764">


Then switch to the branch which you commited by mistake at first and delete that commit by writing the following command 

```
git reset --hard HEAD~1
```
<img width="521" alt="Screenshot 2023-12-29 150840" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/7ef2cd26-6504-426d-b48b-c10f282e56c4">

### Filtering Commit History

Commit history can be filtered :

by date  --before / --after 

by message  --grep 

by author --authore 

by filename --filename 


Example of filtering commit by date :

```
git log --after="2023-12-26"
```
<img width="1058" alt="Screenshot 2024-01-02 191129" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/9812d2dd-0e73-403c-990a-baacfe845ed7">


Example of filtering commit by message :

```
git log --grep="added"
```

<img width="1053" alt="Screenshot 2024-01-02 191223" src="https://github.com/PranitRout07/Advance-Concepts-Of-Github/assets/102309095/7de446a9-f0e2-46f0-91f1-39e9b96d831f">

