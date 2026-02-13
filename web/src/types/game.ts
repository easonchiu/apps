export interface GameData {
  id: string;
  name: string;
  description: string;
  icon: string;
  images: string[];
  appStoreUrl: string;
  content: string[];
}

export const games: Record<string, GameData> = {
  'sudoku-crown': {
    id: 'sudoku-crown',
    name: 'Sudoku Crown',
    description: 'Challenge Your Brain with Sudoku Game!',
    icon: 'https://is1-ssl.mzstatic.com/image/thumb/Purple126/v4/e0/60/4e/e0604e25-034b-78e7-3e69-fbabca7266b6/AppIcon-0-0-1x_U007epad-0-0-0-10-0-0-sRGB-0-0-0-GLES2_U002c0-512MB-85-220-0-0.png/340x340bb.png',
    images: [
      'https://is1-ssl.mzstatic.com/image/thumb/PurpleSource116/v4/80/3a/60/803a602e-85e5-1988-b288-06b34e116a0b/e5f7a9c5-57b8-4ee3-a225-b1f5893b71a3_11.jpg/1242x2688bb.png',
      'https://is1-ssl.mzstatic.com/image/thumb/PurpleSource116/v4/14/12/dc/1412dc7e-ad78-5e7b-c42f-cfe5ea9ed785/b74a1f5f-1b4c-4480-86aa-9a2efbfce23a_22.jpg/1242x2688bb.png',
      'https://is1-ssl.mzstatic.com/image/thumb/PurpleSource126/v4/b9/5f/8f/b95f8f2e-a7b9-886a-d223-806c7a2b2be4/32fedb43-2637-4688-92b5-c70a8a2b90cb_33.jpg/1242x2688bb.png',
      'https://is1-ssl.mzstatic.com/image/thumb/PurpleSource116/v4/a6/c4/6f/a6c46f34-b7d2-85e4-c5ca-c248e44c262b/4348711c-38a7-4aa8-803e-ca06da5b7df3_44.jpg/1242x2688bb.png'
    ],
    appStoreUrl: 'https://apps.apple.com/us/app/sudoku-crown/id6450429527',
    content: [
      'Looking for a game that can exercise your brain? Now, we introduce you to the most popular Sudoku game! Sudoku is a logic game that involves filling a 9x9 grid with numbers. It may sound simple, but it requires high intelligence and patience. Our Sudoku game provides multiple difficulty levels, from beginner to expert, and each game can bring you joy and satisfaction.',
      'Our Sudoku game also includes special features such as note-taking, auto-check, and hints, which can help you solve puzzles more easily. In addition, our Sudoku game offers beautiful interfaces and sound effects, bringing you a brand new gaming experience.',
      'Download our Sudoku game and challenge your brain now! It\'s available for free on the App Store!'
    ]
  },
  'block-cuties': {
    id: 'block-cuties',
    name: 'Block Cuties',
    description: 'An Addictive Block Blast Game! Experience Endless Fun with Block Cuties!',
    icon: 'https://is1-ssl.mzstatic.com/image/thumb/Purple126/v4/7d/a8/1c/7da81c8d-b118-deaa-041f-f0321a80fb74/AppIcon-0-0-1x_U007emarketing-0-0-0-6-0-0-sRGB-0-0-0-GLES2_U002c0-512MB-85-220-0-0.png/340x340bb.png',
    images: [
      'https://is1-ssl.mzstatic.com/image/thumb/PurpleSource116/v4/bf/1e/4e/bf1e4ed9-167c-7372-8fba-f8e5b115a486/c0815d9d-df5a-4ef3-9177-d82e185a1a11_11.png/1242x2688bb.png',
      'https://is1-ssl.mzstatic.com/image/thumb/PurpleSource116/v4/89/35/72/89357206-072b-598d-8bab-30172bdb9372/55a2d251-2a2c-4705-9fa7-da782edd4cea_22.png/1242x2688bb.png',
      'https://is1-ssl.mzstatic.com/image/thumb/PurpleSource116/v4/3b/30/79/3b30791e-df15-8a83-d90b-de4a5e5d7d46/bbd54b05-eb3e-4ed8-9aed-fffad018f096_33.png/1242x2688bb.png'
    ],
    appStoreUrl: 'https://apps.apple.com/us/app/block-cuties/id6459511367',
    content: [
      'The perfect fusion of blocks and match-3 games! Simple controls, delightful sound effects, adorable visuals, all come together to bring you a relaxing and enjoyable gaming experience! Block Cuties will keep you hooked!'
    ]
  },
  'digit-merge': {
    id: 'digit-merge',
    name: 'Digit Merge!',
    description: 'Merge Gears and Balls, Challenge Your Number Intelligence!',
    icon: 'https://is1-ssl.mzstatic.com/image/thumb/Purple126/v4/f3/c4/18/f3c418b6-9ed5-5d7b-9c46-11f5215af1d4/AppIcon-0-0-1x_U007emarketing-0-0-0-6-0-0-sRGB-0-0-0-GLES2_U002c0-512MB-85-220-0-0.png/340x340bb.png',
    images: [
      'https://is1-ssl.mzstatic.com/image/thumb/PurpleSource116/v4/8e/82/6b/8e826b31-7cc7-c37f-0e0e-c25195b75531/1f76b56e-09a1-40cc-b0e3-ab5ec0cac20d_1__U00281_U0029.png/1290x2796bb.png',
      'https://is1-ssl.mzstatic.com/image/thumb/PurpleSource126/v4/4d/71/c6/4d71c661-46e5-af5d-7a47-60b30b58ead1/fe0c9f96-ed67-40da-8356-95eeaa6b452d_2__U00281_U0029.png/1290x2796bb.png',
      'https://is1-ssl.mzstatic.com/image/thumb/PurpleSource126/v4/ae/5f/b0/ae5fb0ba-ae4f-ed97-be17-94a08f0482e0/0c18a4be-b05a-4b0d-93a9-03bbaa954448_3__U00281_U0029.png/1290x2796bb.png',
      'https://is1-ssl.mzstatic.com/image/thumb/PurpleSource116/v4/8a/91/6e/8a916e82-ef9a-23db-a366-519d33df6012/d58aa9a6-e373-468c-b3b9-8f1eb6d0b6b6_4__U00281_U0029.png/1290x2796bb.png'
    ],
    appStoreUrl: 'https://apps.apple.com/us/app/digit-merge/id6448780577',
    content: [
      'Are you ready for a fun and challenging number merging game? Look no further, because this game is perfect for you! Based on number merging, you need to merge gears and balls with numbers to obtain larger starting numbers and numbers for collisions. The starting numbers need to collide to obtain the final number, which can be upgraded to reach even higher scores.',
      'The game is easy to play and offers a simple gameplay experience. You need to merge larger starting numbers and then let them collide to obtain the final number, which will be your score. Each merge can bring surprises and challenge your number intelligence.',
      'If you\'re a fan of number games, then this game is a must-try. Download it now and challenge your number intelligence!'
    ]
  }
};
