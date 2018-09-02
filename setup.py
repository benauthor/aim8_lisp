#!/usr/bin/env python

from setuptools import setup

exec(open('aim8/version.py').read())

setup(
    name='aim8',
    version=__version__,
    author='Evan Bender',
    author_email='maccruiskeen@gmail.com',
    packages=['aim8'],
    entry_points={
          'console_scripts': [
              'aim8 = aim8.interpreter:main',
          ]
      },
)
