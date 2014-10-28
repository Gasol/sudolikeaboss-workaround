sudolikeaboss-workaround
========================

The sudolikeaboss workaround server to facilitate the 1Password5 workaround
that can be referenced here:
https://github.com/ravenac95/sudolikeaboss/issues/1

Disclaimer
----------

The workaround works by editing the 1password extension that is installed on
chrome. This is potentially dangerous. Update at your own risk. AgileBits is
apparently aware of the issue. See:
https://discussions.agilebits.com/discussion/comment/151312. However, there's
no telling how long that will actually take. This workaround was created for
the impatient ones out there (like myself).

Installation
------------

Install with homebrew
*********************

::
    
    $ brew tap ravenac95/sudolikeaboss
    $ brew install sudolikeaboss-workaround


Install from source
*******************

::
    
    $ go get github.com/ravenac95/sudolikeaboss-workaround
    $ cp $GOPATH/bin/sudolikeaboss-workaround /usr/local/bin/sudolikeaboss-workaround

Then go to the src directory of ``sudolikeaboss-workaround`` and run the
following commands::
    
    $ cd scripts
    $ ./sudolikeaboss-setup-workaround
