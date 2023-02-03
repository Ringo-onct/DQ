#include <stdio.h>
#include <stdlib.h>
#include <time.h>

struct PLAYER
{
    int hp;
    int mp;
    int str;
    int action;
};

struct MONSTER
{
    int hp;
    int mp;
    int str;
    int action;
};

int main(void)
{
    struct PLAYER   play;
    struct MONSTER  mons;

    while (1)
    {
        //行動(action)の選択
        printf("0:にげる");
        printf("行動の選択>>");
        scanf("%d", play.action);

        //実際の行動
        //プレイヤーの行動の選択
        switch (play.action)
        {
        case 0:
            printf("にげだした。。。\n");
            break;

        default:
            printf("こんらんしている。\n");
            break;
        }

        if (play.action == 0)
        {
            break;
        }


        //モンスターの行動
        printf("モンスターはようすをみている\n");

    }

    return (0);
}
