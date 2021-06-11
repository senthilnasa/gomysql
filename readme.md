CONTENTS OF THIS FILE
---------------------

 * Introduction
 * Installation
 * Starting the Bot
  * Starting the Bot with Drush
  * Starting the Bot with bot_start.php
 * Using the Bot
 * IRC Message Hooks
 * Other IRC Hooks
 * Design Decisions


INTRODUCTION
------------

Current Maintainer: Morbus Iff <morbus@disobey.com>

Druplicon is an IRC bot that has been servicing #drupal, #drupal-support,
and many other IRC channels since 2005, proving itself an invaluable resource.
Originally a Perl Bot::BasicBot::Pluggable application coded by Morbus Iff,
he always wanted to make the official #drupal bot an actual Drupal module.

This is the fruit of these labors. Whilst the needs of Druplicon are driving
the future and design of the module, this is intended as a generic framework
for IRC bots within Drupal, and usage outside of Druplicon is encouraged.


INSTALLATION
------------

The bot.module is not like other Drupal modules and requires a bit more
effort than normal to get going. Unlike a regular Drupal page load, an
IRC bot has to run forever and, for reasons best explained elsewhere, this
entails running the bot through a shell, NOT through web browser access.

1. This module REQUIRES Net_SmartIRC, a PHP class available from PEAR.
   In most cases, you can simply run "pear install Net_SmartIRC".

2. Copy this bot/ directory to your sites/SITENAME/modules directory.

3. Enable the module(s) and then configure them at admin/config/bot.


STARTING THE BOT
----------------

If you have Drush installed, the following commands are available:

  drush bot-start
  drush bot-status
  drush bot-status-reset
  drush bot-stop

Note, however, that bot.module does use your site's URL in some commands, but
Drush doesn't ever know your URL by default. To ensure proper reporting by the
bot, you'll need to either set your $base_url in settings.php, or tell Drush
your URL with "--uri=http://example.com/". If you don't, the bot will report
all URLs as "http://default/". I prefer setting the URL in settings.php.

To start the bot as a background process, use:

  nohup drush bot-start &

Stopping the bot is accomplished with:

  drush bot-stop

If the bot crashes, is killed, Ctrl-C'd, or otherwise improperly interrupted,
the internal connection status will be stuck in a "connected" state, and
Drush will refuse to "bot-start". You can force a status reset with "drush
bot-status-reset".

IF YOU DO NOT HAVE DRUSH INSTALLED, scripts/bot_start.php is a simple wrapper
around Drupal and the IRC network libraries. To run the bot, move to the
scripts directory and issue the following command:

  php bot_start.php --root /path/to/drupal --url http://www.example.com

--root refers to the full path to your Drupal installation directory and
allows you to execute bot_start.php without moving it to the root directory.
--url is required (and is equivalent to Drupal's base URL) to trick Drupal
into thinking that it is being run as through a web browser. It sets all
the required Drupal environment variables.

If you want to run the bot as a background process, try:

  nohup php bot_start.php --root /path/to/drupal --url http://www.example.com &
